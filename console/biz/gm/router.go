package gm

import (
	accounts "console/biz/gm/accounts/view"
	"console/mods/ginx"
)

func InitGMRouter(g *ginx.RouterGroup) {
	{
		accounts.InitAccountRouter(g.Group("账号管理", "account"))
	}
}
