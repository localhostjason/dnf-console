package service

import (
	"console/biz/gm/model"
	logModel "console/biz/log/model"
	"console/mods/game_db"
	"github.com/localhostjason/webserver/db"
)

func GetRechargeTop5() StatTop5Result {
	return StatTop5Result{
		Cera:           getRechargeTop5ByType("cera"),
		CeraPoint:      getRechargeTop5ByType("cera_point"),
		CeraTotal:      GetRechargeTotal("cera"),
		CeraPointTotal: GetRechargeTotal("cera_point"),
	}
}

func getRechargeTop5ByType(typ string) []TopInfo {
	var data []TopInfo
	tx := db.DB.Model(&logModel.RechargeLog{}).Select("sum(number) as total, uid").Group("action, uid").Order("total desc").Limit(5)

	if typ == "cera" {
		tx.Having("action = 1")
	} else {
		tx.Having("action = 2")
	}
	tx.Find(&data)

	result := make([]TopInfo, 0)
	for _, info := range data {
		info.AccountName = getAccountNameByUid(info.Uid)
		result = append(result, info)
	}
	return result
}

func getAccountNameByUid(uid int) string {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	var data model.Accounts
	dbx.Where("UID = ?", uid).First(&data)
	return data.AccountName
}
