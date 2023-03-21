package cmds

import (
	"console/biz/view"
	"console/mods/casbinx"
	"console/mods/game_db"
	"errors"
	"github.com/localhostjason/webserver/db"
	"github.com/localhostjason/webserver/server"
	"github.com/localhostjason/webserver/svc"

	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type MainProc struct {
	singleMode bool
	quit       chan bool
}

func (m *MainProc) Stop() {
	m.quit <- true
}

func NewProc(singleMode bool) *MainProc {
	return &MainProc{singleMode: singleMode, quit: make(chan bool)}
}

func (m *MainProc) Run(svc *svc.Svc) {
	s, err := startServer(m.singleMode)
	if err != nil {
		return
	}

	<-m.quit
	_ = s.Stop()
}

func (m *MainProc) SigHandlers() map[os.Signal]svc.SignalHandlerFunc {
	return map[os.Signal]svc.SignalHandlerFunc{
		syscall.SIGTERM: m.handleSigTerm,
		os.Interrupt:    m.handleSigTerm,
	}
}

func (m *MainProc) handleSigTerm(sig os.Signal) (err error) {
	m.quit <- true
	return errors.New("quit by signal " + sig.String())
}

func startServer(toConsole bool) (*server.Server, error) {

	s, err := server.NewServer()
	if err != nil {
		log.Fatalln("failed to start:", err)
	}

	err = s.SetLogConfig(toConsole)
	if err != nil {
		log.Fatalln("failed to set log:", err)
	}

	if db.DBEnable() {
		if err = db.Connect(); err != nil {
			log.Fatalln(err)
		}
		if err = db.Migrate(); err != nil {
			log.Fatalln(err)
		}
		if err = db.InitData(); err != nil {
			log.Fatalln(err)
		}
	}

	if game_db.DBEnable() {
		if err = game_db.Connect(); err != nil {
			log.Fatalln(err)
		}
	}

	if err = casbinx.NewCasBin().Run(); err != nil {
		log.Fatalln(err)
	}

	//s.SetRecovery(uv.DefaultRecovery(false))
	if err = s.SetRouter(view.SetView); err != nil {
		return nil, err
	}

	err = s.Start()
	if err != nil {
		log.Fatalln(err)
	}

	return s, nil
}
