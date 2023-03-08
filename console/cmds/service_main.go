package cmds

import (
	"console/mods/casbinx"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
	"github.com/localhostjason/webserver/server"
	"github.com/localhostjason/webserver/svc"

	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var SetViewFunc func(r *gin.Engine) error

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
	if LoadGserverApiFunc != nil {
		s.LoadGrpcServerApi(LoadGserverApiFunc)
	}
	if err != nil {
		log.Fatalln("failed to start:", err)
	}

	err = s.SetLogConfig(toConsole)
	if err != nil {
		log.Fatalln("failed to set log:", err)
	}

	if err = db.Connect(); err != nil {
		log.Fatalln(err)
	}
	if err = db.InitData(); err != nil {
		log.Fatalln(err)
	}

	if err = casbinx.NewCasBin().Run(); err != nil {
		log.Fatalln(err)
	}

	//s.SetRecovery(uv.DefaultRecovery(false))
	if SetViewFunc != nil {
		err = s.SetRouter(SetViewFunc)
		if err != nil {
			return nil, err
		}
	}

	err = s.Start()
	if err != nil {
		log.Fatalln(err)
	}

	return s, nil
}
