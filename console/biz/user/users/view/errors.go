package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_USER_INFO_GET    = "E_USER_INFO_GET"
	E_USER_INFO_UPDATE = "E_USER_INFO_UPDATE"
	E_USERS_GET        = "E_USERS_GET"
)

var eMap = map[string]ue.Error{
	E_USER_INFO_GET:    {Code: E_USER_INFO_GET, Desc: "获取个人用户错误", Msg: "%v"},
	E_USER_INFO_UPDATE: {Code: E_USER_INFO_UPDATE, Desc: "更新个人用户错误", Msg: "%v"},
	E_USERS_GET:        {Code: E_USERS_GET, Desc: "获取用户列表错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
