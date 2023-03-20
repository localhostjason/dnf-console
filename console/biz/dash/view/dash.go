package view

import (
	"console/biz/dash/service"
	"github.com/gin-gonic/gin"
)

func getDashStatCount(c *gin.Context) {
	data := service.GetDashStatTotal()
	c.JSON(200, data)
}

func getDashTop5(c *gin.Context) {
	data := service.GetRechargeTop5()
	c.JSON(200, data)
}
