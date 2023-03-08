package service

import (
	"console/biz/gm/accounts/model"
	"console/mods/game_db"
	"github.com/localhostjason/webserver/server/util/uv"
)

func GetAccounts(pi *uv.PagingIn, order *uv.Order) ([]model.Accounts, *uv.PagingOut, error) {
	dbx := game_db.DBPools.Get(model.AccountDbKey)
	tx := dbx.Model(&model.Accounts{})

	var lst = make([]model.Accounts, 0)
	po, err := uv.PagingFind(tx, &lst, pi, order)

	return lst, po, err
}
