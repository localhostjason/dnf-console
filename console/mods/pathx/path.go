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
/*
1. 目的：在exe（程序下）自动生成 一些目录和默认配置文件
2. GetWd 跟 Executable ，但是 GetWd 类型于 pwd 工作目录，在go run 跑必会是 GetWd = Executable
*/
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
