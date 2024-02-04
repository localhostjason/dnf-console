package view

import "github.com/localhostjason/webserver/server/util/ue"

const (
	E_ACCOUNT_GET = "E_ACCOUNT_GET"

	E_ACCOUNT_DELETE          = "E_ACCOUNT_DELETE"
	E_ACCOUNT_CHANGE_PASSWORD = "E_ACCOUNT_CHANGE_PASSWORD"
	E_ACCOUNT_CREARE          = "E_ACCOUNT_CREARE"
	E_ACCOUNT_UPDATE          = "E_ACCOUNT_UPDATE"

	E_ROLES_GET        = "E_ROLES_GET"
	E_ROLES_UPDATE_QP  = "E_ROLES_UPDATE_QP"
	E_ROLES_UPDATE_PVP = "E_ROLES_UPDATE_PVP"

	E_TASKS_GET     = "E_TASKS_GET"
	E_TASKS_UPDATE  = "E_TASKS_UPDATE"
	E_EMAIL_POST    = "E_EMAIL_POST"
	E_GOLD_GET      = "E_GOLD_GET"
	E_RECHARGE_POST = "E_RECHARGE_POST"

	E_RESET_CREATE_CHARAC = "E_RESET_CREATE_CHARAC"

	E_OPEN_AUCTION = "E_OPEN_AUCTION"
)

var eMap = map[string]ue.Error{
	E_ACCOUNT_GET:             {Code: E_ACCOUNT_GET, Desc: "获取账号错误", Msg: "%v"},
	E_ACCOUNT_DELETE:          {Code: E_ACCOUNT_GET, Desc: "删除账号错误", Msg: "%v"},
	E_ACCOUNT_CHANGE_PASSWORD: {Code: E_ACCOUNT_CHANGE_PASSWORD, Desc: "修改账号密码错误", Msg: "%v"},
	E_ACCOUNT_CREARE:          {Code: E_ACCOUNT_CREARE, Desc: "创建账号错误", Msg: "%v"},
	E_ACCOUNT_UPDATE:          {Code: E_ACCOUNT_UPDATE, Desc: "修改账号错误", Msg: "%v"},

	E_ROLES_GET:        {Code: E_ROLES_GET, Desc: "获取角色错误", Msg: "%v"},
	E_ROLES_UPDATE_QP:  {Code: E_ROLES_UPDATE_QP, Desc: "修改QP错误", Msg: "%v"},
	E_ROLES_UPDATE_PVP: {Code: E_ROLES_UPDATE_PVP, Desc: "修改PVP错误", Msg: "%v"},

	E_TASKS_GET:           {Code: E_TASKS_GET, Desc: "获取任务错误", Msg: "%v"},
	E_TASKS_UPDATE:        {Code: E_TASKS_UPDATE, Desc: "更新任务错误", Msg: "%v"},
	E_EMAIL_POST:          {Code: E_EMAIL_POST, Desc: "发送邮件错误", Msg: "%v"},
	E_GOLD_GET:            {Code: E_GOLD_GET, Desc: "获取物品错误", Msg: "%v"},
	E_RECHARGE_POST:       {Code: E_RECHARGE_POST, Desc: "充值错误", Msg: "%v"},
	E_RESET_CREATE_CHARAC: {Code: E_RESET_CREATE_CHARAC, Desc: "重置创建角色错误", Msg: "%v"},

	E_OPEN_AUCTION: {Code: E_OPEN_AUCTION, Desc: "开启拍卖行错误", Msg: "%v"},
}

func init() {
	ue.RegErrors(eMap)
}
