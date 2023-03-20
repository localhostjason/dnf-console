package view

import (
	"console/biz/log/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getRechargeLog(c *gin.Context) {
	var pi, order, q = &uv.PagingIn{}, &uv.Order{}, &service.RechargeQ{}
	uv.PQ(c, pi, order, q)
	lst, po, err := service.GetReChargeLogList(q, pi, order)
	uv.PEIf(E_LOG_RECHARGE, err)

	c.JSON(200, uv.PagedOut(lst, po))
}

func getOperateLog(c *gin.Context) {
	var pi, order, q = &uv.PagingIn{}, &uv.Order{}, &service.OperateQ{}
	uv.PQ(c, pi, order, q)
	lst, po, err := service.GetOperateLogList(q, pi, order)
	uv.PEIf(E_LOG_RECHARGE, err)

	c.JSON(200, uv.PagedOut(lst, po))
}
