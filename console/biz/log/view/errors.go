package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_LOG_RECHARGE = "E_LOG_RECHARGE"
	E_LOG_OPERATE  = "E_LOG_OPERATE"
)

var eMap = map[string]ue.Error{
	E_LOG_RECHARGE: {Code: E_LOG_RECHARGE, Desc: "获取充值记录错误", Msg: "%v"},
	E_LOG_OPERATE:  {Code: E_LOG_OPERATE, Desc: "获取充值记录错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
