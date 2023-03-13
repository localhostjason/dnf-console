package service

import (
	"console/biz/gm/accounts/model"
	"console/mods/game_db"
	"errors"
	"gorm.io/gorm"
)

func GetRoles(uid int) ([]model.CharacInfo, error) {
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)
	err := dbx.First(&model.Accounts{}, uid).Error

	var data = make([]model.CharacInfo, 0)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("该账号不存在")
	}
	dbx.Table("charac_info").Where("m_id = ?", uid).Find(&data)
	return data, nil
}
