package system

import "daemon/mods/config"

const _key = "dnf_service"

type DnfServiceConfig struct {
	StartShell string `json:"start_shell"`
	StopShell  string `json:"stop_shell"`
}

func getDnfServiceConfig() DnfServiceConfig {
	var c DnfServiceConfig

	_ = config.GetConfig(_key, &c)
	return c
}

func init() {
	c := &DnfServiceConfig{
		StartShell: "sh /home/neople/run",
		StopShell:  "sh /home/neople/stop",
	}
	_ = config.RegConfig(_key, c)
}
