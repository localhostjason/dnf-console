package view

import (
	"console/biz/gm/accounts/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getAccounts(c *gin.Context) {
	var pi, order = &uv.PagingIn{}, &uv.Order{}
	uv.PQ(c, pi, order)
	lst, po, err := service.GetAccounts(pi, order)
	uv.PEIf(E_ACCOUNT_GET, err)

	c.JSON(200, uv.PagedOut(lst, po))
}
