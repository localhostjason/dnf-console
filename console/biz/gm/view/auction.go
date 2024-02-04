package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getAuctionSate(c *gin.Context) {
	data := service.GetAuctionState()
	c.JSON(200, data)
}

func openAuctionSate(c *gin.Context) {
	name := c.Param("name")
	err := service.OpenAuction(name)
	uv.PEIf(E_OPEN_AUCTION, err)
	c.Status(201)
}

func closeAuction(c *gin.Context) {
	c.JSON(200, "未实现")
}
