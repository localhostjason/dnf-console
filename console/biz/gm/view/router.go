package view

import "console/mods/ginx"

func InitGmRouter(r *ginx.RouterGroup) {

	accounts := r.Group("账号相关", "account")
	{
		accounts.GET("获取账号列表", "list", getAccounts)
		accounts.GET("获取角色信息", ":id/roles", getRoles)

		accounts.DELETE("删除账号", ":id", deleteAccount)
		accounts.POST("修改账号密码", ":id/change_password", changePassword)
		accounts.POST("创建账号", "", createAccount)
		accounts.PUT("修改账号信息", ":id", updateAccount)

		accounts.POST("充值", ":id/recharge", rechargeAccount)
		accounts.POST("重置创建角色", ":id/reset_create_charac", resetCreateCharac)
	}

	roles := r.Group("角色管理", "roles")
	{
		roles.PUT("修改QP", ":id/qp", changeQp)
		roles.PUT("修改pk段位", ":id/pvp", changePvp)

		roles.GET("获取任务列表", ":id/tasks", getRoleTasks)
		roles.PUT("更新任务列表", ":id/tasks", updateRoleTasks)

		roles.POST("发送邮件", ":id/email", sendEmail) // id === charac no
	}

	gold := r.Group("物品", "gold")
	{
		gold.GET("获取物品代码", "list", getGolds)
	}

}
