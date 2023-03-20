package service

import (
	"console/biz/gm/model"
	logModel "console/biz/log/model"
	"console/mods/game_db"
	"github.com/localhostjason/webserver/db"
)

func GetDashStatTotal() StatCountResult {
	return StatCountResult{
		CeraTotal:      GetRechargeTotal("cera"),
		CeraPointTotal: GetRechargeTotal("cera_point"),
		UserTotal:      getAccountTotal(),
		CharacTotal:    getCharacTotal(),
	}
}

func GetRechargeTotal(typ string) int {
	type _R struct {
		Total int `json:"total"`
	}
	var data _R
	tx := db.DB.Model(&logModel.RechargeLog{}).Select("sum(number) as total").Group("action")
	if typ == "cera" {
		tx.Having("action = 1").First(&data)
	} else {
		tx.Having("action = 2").First(&data)
	}

	return data.Total
}

func getAccountTotal() int {
	dbx := game_db.DBPools.Get(model.DTaiwan)

	var total int64
	dbx.Model(&model.Accounts{}).Count(&total)
	return int(total)
}

func getCharacTotal() int {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	var accounts []model.Accounts
	dbx.Find(&accounts)

	uids := make([]int, 0)
	for _, info := range accounts {
		uids = append(uids, info.Uid)
	}

	db2 := game_db.DBPools.Get(model.TaiwanCain)

	var total int64
	db2.Table("charac_info").Where("m_id IN ? AND delete_flag = ?", uids, 0).Count(&total)
	return int(total)
}
