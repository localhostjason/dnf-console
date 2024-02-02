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

	var data = make([]model.CharacInfoV2, 0)
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	// 乱码 角色，每次执行一次，无所谓了不搞优化！！！
	//dbx.Exec("UPDATE charac_info SET charac_name = CONVERT(BINARY(CONVERT(charac_name USING latin1)) USING utf8);")
	//
	dbx.Debug().Table("charac_info").Select("CONVERT(BINARY(CONVERT(charac_name USING latin1)) USING utf8) AS converted_charac_name, m_id, charac_no, lev, create_time").Where("m_id = ? AND delete_flag = ?", uid, 0).Find(&data)
	//dbx.Table("charac_info").Where("m_id = ? AND delete_flag = ?", uid, 0).Find(&data)

	var result = make([]RoleResult, 0)
	for _, info := range data {
		result = append(result, RoleResult{
			CharacInfo: model.CharacInfo{
				MId:        info.MId,
				CharacNo:   info.CharacNo,
				CharacName: info.ConvertedCharacName,
				Lev:        info.Lev,
				CreateTime: info.CreateTime,
			},
			Money: getMoneyByRole(info.CharacNo),
			QP:    getQpByRole(info.CharacNo),
			Pvp:   getPvpByRole(info.CharacNo),
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

func getQpByRole(characNo int) int {
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	type _R struct {
		QP int `json:"qp"`
	}
	var data _R
	dbx.Table("charac_quest_shop").Where("charac_no = ?", characNo).Take(&data)
	return data.QP
}

func getPvpByRole(characNo int) PvpResult {
	dbx := game_db.DBPools.Get(model.TaiwanCain)

	var data PvpResult
	dbx.Table("pvp_result").Where("charac_no = ?", characNo).Take(&data)
	return data
}

func UpdateQp(characNo int, args *UpdateQpReq) error {
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	return dbx.Table("charac_quest_shop").Where("charac_no = ?", characNo).Updates(map[string]interface{}{
		"qp": args.QP,
	}).Error
}

func UpdatePvp(characNo int, args *UpdatePvpReq) error {
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	return dbx.Table("pvp_result").Where("charac_no = ?", characNo).Updates(map[string]interface{}{
		"win":       args.Win,
		"pvp_grade": args.PvpGrade,
		"pvp_point": args.PvpPoint,
	}).Error
}
