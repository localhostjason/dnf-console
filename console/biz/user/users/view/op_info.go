package view

import (
	"console/biz/middleware"
	"github.com/localhostjason/webserver/server/util/ue"
)

const (
	I_OP = "I_OP"
)

var iMap = map[string]ue.Info{
	I_OP: {Code: I_OP, Msg: "测试ID: %v, 测试名称：%v"},
}

func init() {
	ue.RegInfos(iMap)
}

const operateKey = middleware.OperateKey
