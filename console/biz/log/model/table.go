package model

import (
	"github.com/localhostjason/webserver/db"
	"time"
)

const (
	RechargeCera      = 1 // action 充值D币
	RechargeCeraPoint = 2 // action 充值D点
)

// RechargeLog 充值記錄
type RechargeLog struct {
	Id     int       `json:"id" gorm:"primaryKey"`
	Uid    int       `json:"uid" gorm:"uid" gorm:"not null"`
	Ip     string    `json:"ip" gorm:"type=string;size=64"`
	Time   time.Time `json:"time" gorm:"not null"`
	Action int       `json:"action"`
	Number int       `json:"number"`
}

// 操作日志

type OperateLog struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Time     time.Time `json:"time" gorm:"not null"`
	Ip       string    `json:"ip" gorm:"type=string;size=64"`
	Action   string    `json:"action"`
	Result   string    `json:"result"`
}

func init() {
	db.RegTables(&RechargeLog{}, &OperateLog{})
}
