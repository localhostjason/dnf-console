package view

import (
	"console/biz/gm/accounts/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getRoles(c *gin.Context) {
	mid := uv.PPID(c, "id")
	roles, err := service.GetRoles(mid)
	uv.PEIf(E_ROLES_GET, err)
	c.JSON(200, roles)
}
