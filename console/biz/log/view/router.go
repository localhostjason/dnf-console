package view

import "console/mods/ginx"

func InitLogRouter(r *ginx.RouterGroup) {
	r.GET("充值日志", "recharge", getRechargeLog)
	r.GET("操作日志", "operate", getOperateLog)
}
