package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_ACCOUNT_GET = "E_ACCOUNT_GET"
	E_ROLES_GET   = "E_ROLES_GET"
)

var eMap = map[string]ue.Error{
	E_ACCOUNT_GET: {Code: E_ACCOUNT_GET, Desc: "获取账号错误", Msg: "%v"},
	E_ROLES_GET:   {Code: E_ROLES_GET, Desc: "获取角色错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
