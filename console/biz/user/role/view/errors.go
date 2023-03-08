package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_USER_ROLE_GET    = "E_USER_ROLE_GET"
	E_USER_ROLE_UPDATE = "E_USER_ROLE_UPDATE"
	E_USER_ROLE_DELETE = "E_USER_ROLE_DELETE"
	E_USER_ROLE_CREATE = "E_USER_ROLE_CREATE"
)

var eMap = map[string]ue.Error{
	E_USER_ROLE_GET:    {Code: E_USER_ROLE_GET, Desc: "获取权限错误", Msg: "%v"},
	E_USER_ROLE_UPDATE: {Code: E_USER_ROLE_UPDATE, Desc: "更新权限错误", Msg: "%v"},
	E_USER_ROLE_DELETE: {Code: E_USER_ROLE_DELETE, Desc: "删除权限错误", Msg: "%v"},
	E_USER_ROLE_CREATE: {Code: E_USER_ROLE_CREATE, Desc: "创建权限错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
