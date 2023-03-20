package model

import (
	"console/biz/user/auth/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
	"time"
)

func CreateRechargeLog(uid int, action int, number int, c *gin.Context) {
	log := &RechargeLog{
		Uid:    uid,
		Time:   time.Now(),
		Action: action,
		Number: number,
		Ip:     c.RemoteIP(),
	}

	db.DB.Create(log)
	return
}

func CreateOperateLog(action string, msg string, c *gin.Context) {
	currentUser := service.CurrentUser(c)
	log := &OperateLog{
		Username: currentUser.Username,
		Time:     time.Now(),
		Action:   action,
		Result:   msg,
		Ip:       c.RemoteIP(),
	}
	db.DB.Create(&log)
	return
}
