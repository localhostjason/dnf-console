package service

import (
	"github.com/localhostjason/webserver/server/util/ue"
)

const (
	I_OP_LOGIN  = "I_OP_LOGIN"
	I_OP_LOGOUT = "I_OP_LOGOUT"
)

var iMap = map[string]ue.Info{
	I_OP_LOGIN:  {Code: I_OP_LOGIN, Action: "登录", Msg: "用户名: %v，登录系统"},
	I_OP_LOGOUT: {Code: I_OP_LOGOUT, Action: "退出登录", Msg: "用户名: %v"},
}

func init() {
	ue.RegInfos(iMap)
}

const operateKey = "_OpLog"
