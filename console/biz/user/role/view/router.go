package view

import "console/mods/ginx"

func InitRoleRouter(r *ginx.RouterGroup) {
	r.GET("获取权限", "", getRole)
	r.DELETE("删除权限", ":id", deleteRole)
	r.PUT("更新权限", ":id", updateRole)
	r.POST("创建权限", "", createRole)
}
