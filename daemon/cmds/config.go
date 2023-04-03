package cmds

import (
	"daemon/mods/config"
	"daemon/mods/pathx"
	"daemon/mods/util"
	"fmt"
	"os"
	"path/filepath"
)

func DumpDefaultConfig() {
	content, err := config.GeneDefaultConfig()
	if err != nil {
		fmt.Println("failed to generate default config")
	} else {
		fmt.Println(string(content))
	}
}

func InitDefaultServerConfigFile(configFile string) string {
	if configFile != "" {
		return configFile
	}

	execPath, _ := pathx.GetExeDir()
	configDir := filepath.Join(execPath, "config")
	logDir := filepath.Join(execPath, "log")
	file := filepath.Join(configDir, "agent.json")

	if !util.PathExists(configDir) {
		_ = os.MkdirAll(configDir, os.ModePerm)
	}

	if !util.PathExists(logDir) {
		_ = os.MkdirAll(logDir, os.ModePerm)
	}

	if util.PathExists(file) {
		return file
	}

	content, _ := config.GeneDefaultConfig()
	_ = os.WriteFile(file, content, os.ModePerm)
	return file
}
