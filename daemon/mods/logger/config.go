package logger

import (
	"daemon/mods/config"
	"daemon/mods/pathx"
	"path/filepath"
)

const _key = "log"

type LogeConfig struct {
	// log
	LogLevel string `json:"log_level"`
	LogPath  string `json:"log_path"`
	ErrorLog string `json:"error_log"`
	SysLog   string `json:"sys_log"`
}

func GetConfig() (LogeConfig, error) {
	var c LogeConfig
	err := config.GetConfig(_key, &c)
	return c, err
}

func init() {
	execPath, _ := pathx.GetExeDir()
	logPath := filepath.Join(execPath, "log")

	c := LogeConfig{
		LogLevel: "info",
		LogPath:  logPath,
		ErrorLog: "error-%Y%m%d.log",
		SysLog:   "sys-%Y%m%d.log",
	}
	_ = config.RegConfig(_key, c)
}
