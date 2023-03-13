package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"errors"
	"gorm.io/gorm"
)

func GetRoles(uid int) ([]model.CharacInfo, error) {
	db1 := game_db.DBPools.Get(model.AccountDbKey)
	err := db1.First(&model.Accounts{}, uid).Error

	var data = make([]model.CharacInfo, 0)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("该账号不存在")
	}
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)
	dbx.Table("charac_info").Where("m_id = ?", uid).Find(&data)
	return data, nil
}
