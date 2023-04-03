package shell

import (
	"daemon/mods/config"
	"time"
)

const _key = "shell"

type configShell struct {
	Timeout     int `json:"timeout"`
	ToolTimeOut int `json:"tool_timeout"`
}

// GetShellToolTimeout 工具超时值
func GetShellToolTimeout() time.Duration {
	var c configShell
	_ = config.GetConfig(_key, &c)
	return time.Duration(c.ToolTimeOut) * time.Second
}

func getConfig() (configShell, error) {
	var c configShell
	err := config.GetConfig(_key, &c)
	return c, err
}

func init() {
	c := configShell{
		Timeout:     30,
		ToolTimeOut: 180,
	}
	_ = config.RegConfig(_key, c)
}
