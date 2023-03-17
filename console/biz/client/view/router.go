package view

import "console/mods/ginx"

func InitClientRouter(r *ginx.RouterGroup) {
	client := r.Group("客户端登录", "")
	{
		client.POST("登录", "login", login)
	}
}
