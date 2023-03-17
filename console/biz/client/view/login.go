package view

import (
	"console/biz/client/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func login(c *gin.Context) {
	data := &service.LoginReq{}
	uv.PB(c, data)

	token, err := service.Login(data)
	uv.PEIf(E_CLIENT_LOGIN, err)
	c.JSON(200, map[string]string{"token": token})
}
