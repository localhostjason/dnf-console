package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetTaskByRole(characNo int) ([]model.Task, error) {
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)

	var characInfo model.CharacInfo
	err := dbx.Table("charac_info").Where("charac_no = ?", characNo).First(&characInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("角色不存在！")
	}

	var data = make([]model.Task, 0)
	dbx.Table("new_charac_quest").Where("charac_no = ?", characNo).Find(&data)
	return data, nil
}

func UpdateTaskByRole(characNo int, ids []int) error {
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)

	var characInfo model.CharacInfo
	err := dbx.Table("charac_info").Where("charac_no = ?", characNo).First(&characInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在！")
	}

	var updateMap = make(map[string]interface{}, 0)
	for i := 1; i <= 20; i++ {
		key := fmt.Sprintf("play_%d_trigger", i)
		updateMap[key] = 0
	}

	fmt.Println(111, characNo, ids, updateMap)

	if len(ids) == 0 {
		return errors.New("请选择任务！")
	}

	for _, id := range ids {
		err = dbx.Debug().Table("new_charac_quest").Where("charac_no = ? AND play_1 = ?", characNo, id).Updates(updateMap).Error
		if err != nil {
			return err
		}
	}
	return nil
}
