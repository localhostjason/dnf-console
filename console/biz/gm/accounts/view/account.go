package view

import (
	"console/biz/gm/accounts/service"
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
