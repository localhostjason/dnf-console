package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getRoles(c *gin.Context) {
	mid := uv.PPID(c, "id")
	roles, err := service.GetRoles(mid)
	uv.PEIf(E_ROLES_GET, err)
	c.JSON(200, roles)
}

func changeQp(c *gin.Context) {
	args := &service.UpdateQpReq{}
	uv.PB(c, args)

	characNo := uv.PPID(c, "id")
	err := service.UpdateQp(characNo, args)
	uv.PEIf(E_ROLES_UPDATE_QP, err)
	c.Status(201)
}
