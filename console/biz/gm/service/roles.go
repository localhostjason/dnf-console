package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"errors"
	"gorm.io/gorm"
)

func GetRoles(uid int) ([]RoleResult, error) {
	db1 := game_db.DBPools.Get(model.DTaiwan)
	err := db1.First(&model.Accounts{}, uid).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该账号不存在")
	}

	var data = make([]model.CharacInfo, 0)
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	dbx.Table("charac_info").Where("m_id = ? AND delete_flag = ?", uid, 0).Find(&data)

	var result = make([]RoleResult, 0)
	for _, info := range data {
		result = append(result, RoleResult{
			CharacInfo: info,
			Money:      getMoneyByRole(info.CharacNo),
		})
	}

	return result, nil
}

func getMoneyByRole(characNo int) int {
	dbx := game_db.DBPools.Get(model.TaiwanCain2nd)

	type _R struct {
		Money int
	}

	var data _R
	dbx.Table("inventory").Where("charac_no = ?", characNo).Take(&data)
	return data.Money
}
