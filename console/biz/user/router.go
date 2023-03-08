package user

import (
	role "console/biz/user/role/view"
	users "console/biz/user/users/view"
	"console/mods/ginx"
)

func InitUserRouter(g *ginx.RouterGroup) {
	{
		role.InitRoleRouter(g.Group("用户权限", "role"))
		users.InitUsersRouter(g.Group("用户信息", ""))
	}
}
