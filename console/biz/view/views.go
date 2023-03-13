package view

import (
	gm "console/biz/gm/view"
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

	apiAuth := r.Group("api/auth")
	api := r.Group("api")

	err = auth.AddJwtAuth(apiAuth, api)
	if err != nil {
		return err
	}

	// load casbin
	api.Use(middleware.CasbinHandler, middleware.ErrorHandler, middleware.OperateHandler)
	routeGroup := ginx.NewRouterGroup(api)
	{
		gm.InitGmRouter(routeGroup.Group("GM管理", "gm"))

		user.InitUserRouter(routeGroup.Group("用户管理", "user"))
	}

	return nil
}
