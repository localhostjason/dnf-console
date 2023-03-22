package pathx

import (
	"os"
	"path/filepath"
)

func getWorkDir() (string, error) {
	// 工作目录 类型 pwd
	exeDir, err := os.Getwd()
	return exeDir, err
}

// GetExeDir DEBUG=1, 将会跑 工作目录
func GetExeDir() (string, error) {
	// 程序目录， exe 二进制文件路径
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)

	if os.Getenv("DEBUG") != "" {
		return getWorkDir()
	}
	return exPath, nil
}
