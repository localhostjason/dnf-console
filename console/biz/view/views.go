package view

import (
	client "console/biz/client/view"
	gm "console/biz/gm/view"
	log "console/biz/log/view"
	"console/biz/middleware"
	"console/biz/static"
	"console/biz/user"
	auth "console/biz/user/auth/view"
	"console/mods/ginx"
	"github.com/gin-gonic/gin"

	_ "console/biz/user/users/service"
)

func SetView(r *gin.Engine) error {
	c, err := GetConfig()
	if err != nil {
		return err
	}
	if c.CORS {
		setCORS(r)
	}

	static.AddStaticToRouter(r)
	r.Use(middleware.OperateHandler)

	apiAuth := r.Group("api/auth")
	api := r.Group("api")

	err = auth.AddJwtAuth(apiAuth, api)
	if err != nil {
		return err
	}

	// load casbin
	api.Use(middleware.CasbinHandler, middleware.ErrorHandler)
	routeGroup := ginx.NewRouterGroup(api)
	{
		gm.InitGmRouter(routeGroup.Group("GM管理", "gm"))
		client.InitClientRouter(routeGroup.Group("客户端", "client"))

		user.InitUserRouter(routeGroup.Group("用户管理", "user"))
		log.InitLogRouter(routeGroup.Group("日志", "log"))
	}

	return nil
}
