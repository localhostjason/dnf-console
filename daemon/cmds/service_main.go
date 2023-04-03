package cmds

import (
	"daemon/biz/system"
	"daemon/mods/logger"
	"daemon/mods/svc"
	"errors"
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type MainProc struct {
	singleMode bool
	quit       chan bool

	DnfService *system.DnfService
}

func (m *MainProc) Stop() {
	m.DnfService.StopChan()
	m.quit <- true
}

func NewProc(singleMode bool) *MainProc {
	return &MainProc{singleMode: singleMode, quit: make(chan bool)}
}

func (m *MainProc) Run(svc *svc.Svc) {
	if err := startAgent(true); err != nil {
		return
	}

	dnfService := system.NewDnfService()
	if err := dnfService.Run(); err != nil {
		log.Fatalln(err)
	}
	m.DnfService = dnfService
	<-m.quit
	dnfService.Stop()
}

func (m *MainProc) SigHandlers() map[os.Signal]svc.SignalHandlerFunc {
	return map[os.Signal]svc.SignalHandlerFunc{
		syscall.SIGTERM: m.handleSigTerm,
		os.Interrupt:    m.handleSigTerm,
	}
}

func (m *MainProc) handleSigTerm(sig os.Signal) (err error) {
	m.DnfService.StopChan()
	m.quit <- true
	return errors.New("quit by signal " + sig.String())
}

func startAgent(toConsole bool) error {
	err := logger.SetLogConfig(toConsole)
	if err != nil {
		log.Fatalln("failed to set log:", err)
	}
	return nil
}
