package service

import (
	"console/biz/gm/accounts/model"
	"console/mods/game_db"
	"errors"
	"github.com/localhostjason/webserver/server/util/uv"
	"gorm.io/gorm"
)

func GetAccounts(pi *uv.PagingIn, order *uv.Order) ([]AccountResult, *uv.PagingOut, error) {
	dbx := game_db.DBPools.Get(model.AccountDbKey)
	tx := dbx.Model(&model.Accounts{})

	var lst = make([]model.Accounts, 0)
	po, err := uv.PagingFind(tx, &lst, pi, order)

	var data = make([]AccountResult, 0)
	for _, info := range lst {
		roles := getGameRolesByUid(info.Uid)

		money := getGameMoneyByUid(info.Uid)

		data = append(data, AccountResult{
			Accounts: info,
			Roles:    roles,
			Money:    money,
		})
	}

	return data, po, err
}

func getGameRolesByUid(uid int) int64 {
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)

	var total int64
	dbx.Table("charac_info").Where("m_id = ?", uid).Count(&total)
	return total
}

func getGameMoneyByUid(uid int) any {
	dbx := game_db.DBPools.Get(model.AccountDetailDbKey)
	type _R struct {
		Money    int `json:"money"`
		Capacity int `json:"capacity"`
	}

	result := map[string]interface{}{}
	err := dbx.Debug().Table("account_cargo").Where("m_id = ?", uid).Take(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}
	return result["money"]
}
