package pathx

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExeDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res := filepath.Dir(exePath)

	if strings.Contains(exePath, getTmpDir()) {
		// run 模式下，确在当前程序入口目录
		return os.Getwd()
	}
	return res, nil
}

func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	return dir
}
