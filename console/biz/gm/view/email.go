package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func sendEmail(c *gin.Context) {
	characNo := uv.PPID(c, "id")

	email := &service.Email{}
	uv.PB(c, email)
	err := service.SendEmail(characNo, email)
	uv.PEIf(E_EMAIL_POST, err)
	c.Status(201)
}
