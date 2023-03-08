package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const OperateKey = "_OpLog"

func OperateHandler(c *gin.Context) {
	c.Next()

	// 全局操作 入庫
	data, ok := c.Get(OperateKey)
	if ok {
		fmt.Println("operate:", data)
	}

}
