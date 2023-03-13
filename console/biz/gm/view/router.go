package view

import "console/mods/ginx"

func InitGmRouter(r *ginx.RouterGroup) {

	accounts := r.Group("账号相关", "account")
	{
		accounts.GET("获取账号列表", "list", getAccounts)
		accounts.GET("获取角色信息", ":id/roles", getRoles)
	}

	tasks := r.Group("任务清理", "roles")
	{
		tasks.GET("获取任务列表", ":id/tasks", getRoleTasks)
		tasks.PUT("更新任务列表", ":id/tasks", updateRoleTasks)

		tasks.POST("发送邮件", ":id/email", sendEmail) // id === charac no
	}

	gold := r.Group("物品", "gold")
	{
		gold.GET("获取物品代码", "list", getGolds)
	}

}
