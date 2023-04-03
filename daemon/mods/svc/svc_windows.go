package svc

import (
	"flag"
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
	"os/signal"
	"time"
)

type Svc struct {
	name        string
	description string
	task        Task
	singleMode  bool
}

func NewSvc(name, description string, task Task) *Svc {
	return &Svc{
		name:        name,
		description: description,
		task:        task,
	}
}

func (s *Svc) setupSigHandlers() chan os.Signal {
	// signals, from github.com/sevlyar/go-daemon
	sigHandlers := s.task.SigHandlers()
	signals := make([]os.Signal, 0, len(sigHandlers))
	for sig := range sigHandlers {
		signals = append(signals, sig)
	}
	ch := make(chan os.Signal, 8)
	signal.Notify(ch, signals...)
	return ch
}

func (s *Svc) Install() {
	m := s.connectMgr()
	defer m.Disconnect()

	ss, err := m.OpenService(s.name)
	if err == nil {
		ss.Close()
		panicErr("service already exists " + ss.Name)
	}

	exePath := exePath()
	exeArg := []string{"-k", cmdService}

	ss, err = m.CreateService(s.name, exePath,
		mgr.Config{StartType: mgr.StartAutomatic, DisplayName: s.name, Description: s.description}, exeArg...)
	if err != nil {
		panicErr("failed to create service " + err.Error())
	}
}
func (s *Svc) open(m *mgr.Mgr) *mgr.Service {
	ss, err := m.OpenService(s.name)
	if err != nil {
		panicErr("could not open service " + err.Error())
	}
	return ss
}

func (s *Svc) openService() (*mgr.Mgr, *mgr.Service) {
	m := s.connectMgr()
	ss := s.open(m)
	return m, ss
}

func (s *Svc) connectMgr() *mgr.Mgr {
	m, err := mgr.Connect()
	if err != nil {
		panicErr("failed to connect service manager " + err.Error())
	}
	return m
}

func (s *Svc) Uninstall() {
	m, ss := s.openService()
	defer m.Disconnect()
	defer ss.Close()

	err := ss.Delete()
	if err != nil {
		panicErr("failed to delete service " + s.name + " " + err.Error())
	}
}

// should quit when return true
func (s *Svc) handleSig(sig os.Signal) bool {
	return s.task.SigHandlers()[sig](sig) != nil
}

func (s *Svc) RunSingle() {
	s.singleMode = true
	log.Println("run single mode")
	ch := s.setupSigHandlers()
	go s.task.Run(s)

loop:
	for {
		select {
		case sig := <-ch:
			log.Println("got signal", sig)
			if s.handleSig(sig) {
				log.Println("quit by signal", sig)
				signal.Stop(ch)
				break loop
			}
		}
	}
}

// RunService 服务自动实际运行的业务代码. Start/Stop()只是发出一个调度通知
func (s *Svc) RunService() bool {
	err := svc.Run(s.name, s)
	if err != nil {
		panicErr("failed to run service " + err.Error())
	}
	return true
}

func (s *Svc) Start() {
	m, ss := s.openService()
	defer m.Disconnect()
	defer ss.Close()

	err := ss.Start()
	if err != nil {
		panicErr("could not start service " + err.Error())
	}
}

func (s *Svc) Run() {
	s.Start()
}

func (s *Svc) Stop() {
	m, ss := s.openService()
	defer m.Disconnect()
	defer ss.Close()

	status, err := ss.Control(svc.Stop)
	if err != nil {
		panicErr("could not send stop to service " + err.Error())
	}

	timeout := time.Now().Add(15 * time.Second)
	for status.State != svc.Stopped {
		if timeout.Before(time.Now()) {
			panicErr("timeout waiting for service to stop " + ss.Name)
		}
		time.Sleep(300 * time.Millisecond)
		status, err = ss.Query()
		if err != nil {
			panicErr("could not retrieve service status " + err.Error())
		}
	}
}

func (s *Svc) SingleMode() bool {
	return s.singleMode
}

// Status running or not
func (s *Svc) Status() (e error) {
	defer func() {
		err := recover()
		if err != nil {
			e = err.(error)
		}
	}()
	m := s.connectMgr()
	defer m.Disconnect()

	ss, err := m.OpenService(s.name)
	if err != nil {
		panicErr("could not open service " + err.Error())
	}
	defer ss.Close()
	state, err := ss.Query()
	if err != nil {
		panicErr("failed to query service status " + err.Error())
	}
	if state.State != svc.Running {
		panicErr(fmt.Sprintln("not running, state:", state.State))
	}
	return nil
}

func (s *Svc) Execute(
	args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	//signals, from github.com/sevlyar/go-daemon
	ch := s.setupSigHandlers()

	go s.task.Run(s)

loop:
	for {
		select {
		// service schedule ...
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				// golang.org/x/sys/windows/svc.TestExample is verifying this output.
				changes <- svc.Status{State: svc.StopPending}
				s.task.Stop()
				break loop
			default:
				log.Println("unexpected control request #", c)
			}
		//signal handling
		case sig := <-ch:
			log.Println("got signal", sig)
			if s.handleSig(sig) {
				log.Println("quit by signal handler")
				signal.Stop(ch)
				break loop
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func (s *Svc) DefaultRun() {
	singleMode := flag.Bool("x", false, "start, no service mode")
	cmd := flag.String("k", "", "`cmds`: start|stop|status|install|uninstall")
	flag.Parse()

	switch *cmd {
	case cmdInstall:
		s.Install()
	case cmdUninstall:
		s.Uninstall()
	case cmdStart:
		s.Start()
	case cmdStop:
		s.Stop()
	case cmdStatus:
		if s.Status() == nil {
			fmt.Println("running")
		} else {
			fmt.Println("not running")
		}
	case cmdService:
		s.RunService()
	default:
		if *singleMode {
			s.RunSingle()
		} else {
			flag.PrintDefaults()
		}
	}
}

func (s *Svc) RunMain(singleMode bool, cmd string) {
	switch cmd {
	case cmdInstall:
		s.Install()
	case cmdUninstall:
		s.Uninstall()
	case cmdStart:
		s.Start()
	case cmdStop:
		s.Stop()
	case cmdStatus:
		if s.Status() == nil {
			fmt.Println("running")
		} else {
			fmt.Println("not running")
		}
	case cmdService:
		s.RunService()
	default:
		if singleMode {
			s.RunSingle()
		} else {
			flag.PrintDefaults()
		}
	}
}
