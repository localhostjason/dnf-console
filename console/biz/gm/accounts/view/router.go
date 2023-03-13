package view

import "console/mods/ginx"

func InitAccountRouter(r *ginx.RouterGroup) {

	accounts := r.Group("账号管理", "list")
	{
		accounts.GET("获取账号列表", "", getAccounts)
	}

	roles := r.Group("角色管理", ":id")
	{
		roles.GET("获取角色信息", "roles", getRoles)
	}

}
