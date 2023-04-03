package cmds

import (
	"daemon/mods/config"
	"daemon/mods/pathx"
	"daemon/mods/svc"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const _groupService = "service"

type ServiceConfig struct {
	LogPath   string `json:"log_path"`
	PidFile   string `json:"pid_file"`
	DaemonLog string `json:"daemon_log"`
}

func getConf() (ServiceConfig, error) {
	var c ServiceConfig
	err := config.GetConfig(_groupService, &c)
	if err != nil {
		return c, err
	}
	c.PidFile = filepath.Join(c.LogPath, c.PidFile)
	c.DaemonLog = filepath.Join(c.LogPath, c.DaemonLog)
	if err := os.MkdirAll(c.LogPath, 0755); err != nil {
		return c, fmt.Errorf("failed to create log dir %s", c.LogPath)
	}

	return c, nil
}

func init() {
	execPath, _ := pathx.GetExeDir()
	logPath := filepath.Join(execPath, "log")
	c := ServiceConfig{
		LogPath:   logPath,
		PidFile:   "console.pid",
		DaemonLog: "daemon.log",
	}
	_ = config.RegConfig(_groupService, c)
}

func NewService(prc *MainProc) (*svc.Svc, error) {
	if runtime.GOOS == "windows" {
		svcName := "agent"
		svcDescription := "agent 服务"
		return svc.NewSvc(svcName, svcDescription, prc), nil
	} else {
		c, err := getConf()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to get daemon config:%v", err))
		}
		return svc.NewSvc(c.PidFile, c.DaemonLog, prc), nil
	}
}
