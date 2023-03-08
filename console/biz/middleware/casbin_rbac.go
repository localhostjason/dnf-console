package middleware

import (
	roleService "console/biz/user/role/service"
	"console/mods/casbinx"
	"github.com/gin-gonic/gin"
)

func CasbinHandler(c *gin.Context) {
	// 获取请求的URI
	obj := c.Request.URL.RequestURI()
	// 获取请求方法
	act := c.Request.Method
	// 获取用户的角色
	user := roleService.GetCurrentUser(c)
	sub := user.Role

	// 判断策略中是否存在
	success, _ := casbinx.E.Enforce(sub, obj, act)

	if success {
		c.AbortWithStatus(403)
	} else {
		c.Next()
	}

}
