package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getAccounts(c *gin.Context) {
	var pi, order, q = &uv.PagingIn{}, &uv.Order{}, &service.AccountFilter{}
	uv.PQ(c, pi, order, q)
	lst, po, err := service.GetAccounts(q, pi, order)
	uv.PEIf(E_ACCOUNT_GET, err)

	c.JSON(200, uv.PagedOut(lst, po))
}

func rechargeAccount(c *gin.Context) {
	uid := uv.PPID(c, "id")

	data := &service.RechargeReq{}
	uv.PB(c, data)

	err := service.RechargeAccount(uid, data)
	uv.PEIf(E_RECHARGE_POST, err)
	c.Status(201)
}
