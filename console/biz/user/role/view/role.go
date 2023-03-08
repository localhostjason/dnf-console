package view

import (
	"console/biz/user/role/model"
	"console/biz/user/role/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
	"strconv"
)

func createRole(c *gin.Context) {
	rq := &model.RolePost{}
	uv.PB(c, rq)

	err := service.CreateRole(rq.Role, rq.Path, rq.Method)
	uv.PEIf(E_USER_ROLE_CREATE, err)
	c.Status(201)
}

func deleteRole(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := service.DeleteRole(idInt)
	uv.PEIf(E_USER_ROLE_DELETE, err)
	c.JSON(200, service.GetRolePolicy())
}

func updateRole(c *gin.Context) {
	err := service.UpdateRole(1, "admin", "/api/user/password", "GET")
	uv.PEIf(E_USER_ROLE_UPDATE, err)
	c.JSON(200, service.GetRolePolicy())
}

func getRole(c *gin.Context) {
	data := service.GetRolePolicy()
	c.JSON(200, data)
}
