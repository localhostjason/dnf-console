package middleware

import (
	"console/biz/log/model"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/ue"
)

const OperateKey = "_OpLog"

func OperateHandler(c *gin.Context) {
	c.Next()

	// 全局操作 入庫
	data, ok := c.Get(OperateKey)
	if ok {
		info, ok2 := data.(*ue.Info)
		if ok2 {
			model.CreateOperateLog(info.Action, info.Msg, c)
		}
		//fmt.Println(11111, data)
	}

}
