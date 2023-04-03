package config

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	TimeFormat        = "2006-01-02 15:04:05"
	TimeFormatNoBlank = "20060102_150405"
	TimeFormatDay     = "20060102"
)

var RootDir string

func init() {
	RootDir = getExePath2()
}

func getExePath() string {
	ex, _ := os.Executable()
	exeDir, _ := filepath.Abs(filepath.Dir(ex))
	return exeDir
}

func getExePath2() string {
	exeDir, _ := GetExeDir()
	return exeDir
}

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
