package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"errors"
	"github.com/localhostjason/webserver/server/util/uv"
	"gorm.io/gorm"
)

func GetAccounts(q *AccountFilter, pi *uv.PagingIn, order *uv.Order) ([]AccountResult, *uv.PagingOut, error) {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	tx := q.FilterQuery(dbx)

	var lst = make([]model.Accounts, 0)
	po, err := uv.PagingFind(tx, &lst, pi, order)

	var data = make([]AccountResult, 0)
	for _, info := range lst {
		roles := getGameRolesByUid(info.Uid)

		money, capacity := getGameMoneyByUid(info.Uid)
		ceraPoint, cera := getCashByUid(info.Uid)

		data = append(data, AccountResult{
			Accounts:  info,
			Roles:     roles,
			Money:     money,
			Capacity:  capacity,
			CeraPoint: ceraPoint,
			Cera:      cera,
		})
	}

	return data, po, err
}

func getGameRolesByUid(uid int) int64 {
	dbx := game_db.DBPools.Get(model.TaiwanCain)

	var total int64
	dbx.Table("charac_info").Where("m_id = ?", uid).Count(&total)
	return total
}

func getGameMoneyByUid(uid int) (int, int) {
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	type _R struct {
		Money    int `json:"money"`
		Capacity int `json:"capacity"`
	}

	var result _R
	err := dbx.Debug().Table("account_cargo").Where("m_id = ?", uid).Take(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0
	}

	return result.Money, result.Capacity
}

func getCashByUid(uid int) (int, int) {
	dbx := game_db.DBPools.Get(model.TaiwanBilling)

	type CashPoint struct {
		CeraPoint int `gorm:"column:cera_point"`
	}
	var data CashPoint
	dbx.Table("cash_cera_point").Where("account = ?", uid).Take(&data)

	type CashCera struct {
		Cera    int
		ModTran int
	}
	var d CashCera
	dbx.Table("cash_cera").Where("account = ?", uid).Take(&d)

	return data.CeraPoint, d.Cera
}

func RechargeAccount(uid int, data *RechargeReq) error {
	dbx := game_db.DBPools.Get(model.TaiwanBilling)
	if data.CeraOption == "cera" {
		err := dbx.Table("cash_cera").Where("account = ?", uid).Updates(map[string]interface{}{
			"cera": data.Cera,
		}).Error
		return err
	}
	errPoint := dbx.Table("cash_cera_point").Where("account = ?", uid).Updates(map[string]interface{}{
		"cera_point": data.CeraPoint,
	}).Error

	return errPoint
}
