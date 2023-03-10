package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_ACCOUNT_GET  = "E_ACCOUNT_GET"
	E_ROLES_GET    = "E_ROLES_GET"
	E_TASKS_GET    = "E_TASKS_GET"
	E_TASKS_UPDATE = "E_TASKS_UPDATE"
	E_EMAIL_POST   = "E_EMAIL_POST"
)

var eMap = map[string]ue.Error{
	E_ACCOUNT_GET:  {Code: E_ACCOUNT_GET, Desc: "获取账号错误", Msg: "%v"},
	E_ROLES_GET:    {Code: E_ROLES_GET, Desc: "获取角色错误", Msg: "%v"},
	E_TASKS_GET:    {Code: E_TASKS_GET, Desc: "获取任务错误", Msg: "%v"},
	E_TASKS_UPDATE: {Code: E_TASKS_UPDATE, Desc: "更新任务错误", Msg: "%v"},
	E_EMAIL_POST:   {Code: E_EMAIL_POST, Desc: "发送邮件错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
