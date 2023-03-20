package view

import "console/mods/ginx"

func InitDashRouter(r *ginx.RouterGroup) {
	r.GET("获取首页统计", "stat/count", getDashStatCount)
	r.GET("获取首页Top5", "stat/top5", getDashTop5)
	//r.GET("操作日志", "operate", getOperateLog)
}
