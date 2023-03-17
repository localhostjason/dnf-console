package cmds

import (
	"console/mods/game_db"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server"
	"github.com/localhostjason/webserver/server/config"
	"google.golang.org/grpc"
)

type MainWorkFunc func(r *gin.Engine) error

type MainServer struct {
	DefaultConfigPath string
}

func NewMainServer() *MainServer {
	return &MainServer{}
}

func (m *MainServer) SetServerConfigFile(file string) {
	m.DefaultConfigPath = file
}

// Run 可根据自己业务 替换扩展
func (m *MainServer) Run() {
	defaultConfigFile := server.InitDefaultServerConfigFile(m.DefaultConfigPath)

	configPath := flag.String("p", defaultConfigFile, "path to config")
	initDB := flag.Bool("i", false, "int db")
	dumpConfig := flag.Bool("d", false, "dump default config")

	createPem := flag.Bool("pem", false, "create pem")

	// for service
	singleMode := flag.Bool("x", false, "start, no daemon/service mode")
	svcCMD := flag.String("k", "", "cmds:start|stop|status, windows: install|uninstall")

	flag.Parse()

	if err := config.SetConfigFile(*configPath); err != nil {
		fmt.Println("failed to set config path", *configPath, err)
		return
	}

	// commands

	if *dumpConfig {
		DumpDefaultConfig()
		return
	}

	if *createPem {
		if err := CreatePem(); err != nil {
			fmt.Println("error create pem", err)
			return
		}
		fmt.Println("success create publicKey.pem / private.pem")
		return
	}

	// DB 初始表结构和默认值
	if *initDB {
		if err := SyncDB(); err != nil {
			fmt.Println("error when sync db schema", err)
			return
		}
		if err := game_db.SyncDB(); err != nil {
			fmt.Println("error when sync game db schema", err)
			return
		}

		fmt.Println("success: sync db schema")
		return
	}

	if err := AutoMigrate(); err != nil {
		fmt.Println("error when migrate db schema", err)
		return
	}
	if err := game_db.AutoMigrate(); err != nil {
		fmt.Println("error when migrate game db schema", err)
		return
	}

	RunService(*singleMode, *svcCMD)
}

var LoadGserverApiFunc func(server *grpc.Server)

func (m *MainServer) LoadGrpcServerApi(loadFunc func(*grpc.Server)) {
	LoadGserverApiFunc = loadFunc
}

func (m *MainServer) LoadView(setView MainWorkFunc) {
	SetViewFunc = setView
}
