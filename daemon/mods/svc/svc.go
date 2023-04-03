package svc

import "os"

// SignalHandlerFunc 如果返回错误，程序会退出，但不会主动调用Stop()。退出相关事宜由handler自行处理
type SignalHandlerFunc func(sig os.Signal) (err error)

type Task interface {
	// Run 开始任务，不要返回。如果返回，则表示程序结束
	Run(svc *Svc)

	// Stop 通知Task结束Run
	Stop()

	// SigHandlers 信号处理函数
	SigHandlers() map[os.Signal]SignalHandlerFunc
}

const (
	cmdStart     = "start"
	cmdStop      = "stop"
	cmdStatus    = "status"
	cmdInstall   = "install"
	cmdUninstall = "uninstall"
	cmdService   = "service"
)
