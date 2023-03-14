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

func getGolds(c *gin.Context) {
	var pi, order, q = &uv.PagingIn{}, &uv.Order{}, &service.GoldQ{}
	uv.PQ(c, pi, order, q)

	lst, po, err := service.GetGoldList(q, pi, order)
	uv.PEIf(E_GOLD_GET, err)
	c.JSON(200, uv.PagedOut(lst, po))
}
