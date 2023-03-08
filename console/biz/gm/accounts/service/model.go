package service

import (
	"console/biz/gm/accounts/model"
	"github.com/localhostjason/webserver/db"
	"gorm.io/gorm"
)

type AccountResult struct {
	model.Accounts
	Roles int64 `json:"roles"`
	Money any   `json:"money"`
}

type AccountFilter struct {
	Uid string `form:"uid"`
}

func (q AccountFilter) FilterQuery(dbx *gorm.DB) (tx *gorm.DB) {
	tx = dbx.Model(&model.Accounts{})

	if q.Uid != "" {
		tx = tx.Where("UID like ?", db.Like(q.Uid))
	}
	return
}
