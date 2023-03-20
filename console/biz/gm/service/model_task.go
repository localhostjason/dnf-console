package service

import (
	"console/biz/gm/model"
	"github.com/localhostjason/webserver/db"
	"gorm.io/gorm"
)

type TaskId struct {
	Ids []int `json:"ids"`
}

type Email struct {
	Code            int  `json:"code" binding:"gte=0"`
	Number          int  `json:"number" binding:"required"`
	SeperateUpgrade int  `json:"seperate_upgrade"` // 武将锻造等级
	Upgrade         int  `json:"upgrade"`          // 装备等级
	IsAmplify       bool `json:"is_amplify"`       // 具有异界属性
	AmplifyOption   int  `json:"amplify_option"`
	AmplifyValue    int  `json:"amplify_value"`
	Gold            int  `json:"gold"`      // 金币
	SealFlag        bool `json:"seal_flag"` // 是否封装
}

type GoldQ struct {
	Name string `form:"name"`
}

func (q GoldQ) FilterQuery(dbx *gorm.DB) (tx *gorm.DB) {
	tx = dbx.Model(&model.Gold{})

	if q.Name != "" {
		tx = tx.Where("name like ?", db.Like(q.Name))
	}
	return
}
