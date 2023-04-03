package ue

import "fmt"

// 使用时 要先注册一个error map

type Error struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}

func NewErr(code string, args ...interface{}) *Error {
	e, found := _eMap[code]
	if !found {
		panic("webserver.util.ue, invalid error code")
	}
	return &Error{
		Code: code,
		Msg:  fmt.Sprintf(e.Msg, args...),
		Desc: e.Desc,
	}
}

var _eMap map[string]Error

func RegErrors(m map[string]Error) {
	if _eMap == nil {
		_eMap = make(map[string]Error)
	}
	for k, v := range m {
		_eMap[k] = v
	}
}
