package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getRoleTasks(c *gin.Context) {
	characNo := uv.PPID(c, "id")
	data, err := service.GetTaskByRole(characNo)
	uv.PEIf(E_TASKS_GET, err)
	c.JSON(200, data)
}

func updateRoleTasks(c *gin.Context) {
	characNo := uv.PPID(c, "id")

	args := &service.TaskId{}
	uv.PB(c, args)

	err := service.UpdateTaskByRole(characNo, args.Ids)
	uv.PEIf(E_TASKS_UPDATE, err)
	c.Status(201)
}
