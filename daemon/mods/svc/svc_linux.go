package svc

import (
	"errors"
	"flag"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Svc struct {
	pidFile    string
	logFile    string
	task       Task
	singleMode bool
}

func NewSvc(pid_file, log_file string, task Task) *Svc {
	return &Svc{
		pidFile: pid_file,
		logFile: log_file,
		task:    task,
	}
}

// should quit when return true
func (s *Svc) handleSig(sig os.Signal) bool {
	return s.task.SigHandlers()[sig](sig) != nil
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

func (s *Svc) RunSingle() {
	log.Println("run single")
	ch := s.setupSigHandlers()
	s.singleMode = true
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
func (s *Svc) newContext() *daemon.Context {
	return &daemon.Context{
		PidFileName: s.pidFile,
		PidFilePerm: 0644,
		LogFileName: s.logFile,
		LogFilePerm: 0644,
	}
}

func (s *Svc) SingleMode() bool {
	return s.singleMode
}

func (s *Svc) Run() {
	ctx := s.newContext()
	child, err := ctx.Reborn()
	if err != nil {
		if errors.Is(err, daemon.ErrWouldBlock) {
			panicErr("The program may already running, err: " + err.Error())
		} else {
			panicErr("failed to create child process, err: " + err.Error())
		}
	}

	if child != nil {
		// parent
		return
	}
	defer ctx.Release()

	for sig, handler := range s.task.SigHandlers() {
		h := daemon.SignalHandlerFunc(handler)
		daemon.SetSigHandler(h, sig)
	}

	// clear parent prepareEnv
	clearEnv()

	go s.task.Run(s)

	if err := daemon.ServeSignals(); err != nil {
		log.Println("serve signal() end", err)
	}
}

func clearEnv() {
	if _, present := os.LookupEnv(daemon.MARK_NAME); present {
		_ = os.Unsetenv(daemon.MARK_NAME)
	}
}

func (s *Svc) Stop() {
	ctx := s.newContext()
	d, err := ctx.Search()
	if err != nil {
		panicErr("failed to find process " + err.Error())
		return
	}
	err = d.Signal(syscall.SIGTERM)
	if err != nil {
		panicErr(fmt.Sprintf("failed to send term signal to pid %d, %s", d.Pid, err))
	}
}

func (s *Svc) Status() (e error) {
	ctx := s.newContext()
	d, err := ctx.Search()
	if err == nil && d != nil {
		return nil
	} else {
		return errors.New("could not find instance")
	}
}

// 根据命令行程序，执行默认功能
func (s *Svc) DefaultRun() {
	singleMode := flag.Bool("x", false, "start , no daemon or service mode")
	cmd := flag.String("k", "", "`cmds`: start|stop|status")
	flag.Parse()

	switch *cmd {
	case cmdStart:
		s.Run()
	case cmdStop:
		s.Stop()
	case cmdStatus:
		if s.Status() == nil {
			fmt.Println("running")
		} else {
			fmt.Println("not running")
		}
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
	case cmdStart:
		s.Run()
	case cmdStop:
		s.Stop()
	case cmdStatus:
		if s.Status() == nil {
			fmt.Println("running")
		} else {
			fmt.Println("not running")
		}
	default:
		if singleMode {
			s.RunSingle()
		} else {
			flag.PrintDefaults()
		}
	}
}
