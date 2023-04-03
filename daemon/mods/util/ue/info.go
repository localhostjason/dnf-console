package ue

import "fmt"

// 使用时 要先注册一个info map ,给 操作日志用

type Info struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Info) OperateInfo() string {
	return e.Msg
}

func NewInfo(code string, args ...interface{}) *Info {
	e, found := _iMap[code]
	if !found {
		panic("webserver.util.ue, invalid info code")
	}
	return &Info{
		Code: code,
		Msg:  fmt.Sprintf(e.Msg, args...),
	}
}

var _iMap map[string]Info

func RegInfos(m map[string]Info) {
	if _iMap == nil {
		_iMap = make(map[string]Info)
	}
	for k, v := range m {
		_iMap[k] = v
	}
}
