package view

import "console/mods/ginx"

func InitDashRouter(r *ginx.RouterGroup) {
	r.GET("获取首页统计", "count/stat", getDashStatCount)
	//r.GET("操作日志", "operate", getOperateLog)
}
