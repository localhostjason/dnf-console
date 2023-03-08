package view

import "console/mods/ginx"

func InitAccountRouter(r *ginx.RouterGroup) {

	accounts := r.Group("账号管理", "list")
	{
		accounts.GET("获取账号列表", "", getAccounts)

	}

}
