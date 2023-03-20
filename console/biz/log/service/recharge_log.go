package service

import (
	"console/biz/log/model"
	"github.com/localhostjason/webserver/db"
	"github.com/localhostjason/webserver/server/util/uv"
)

func GetReChargeLogList(q *RechargeQ, pi *uv.PagingIn, order *uv.Order) ([]model.RechargeLog, *uv.PagingOut, error) {
	tx := q.FilterQuery(db.DB)
	var lst = make([]model.RechargeLog, 0)

	po, err := uv.PagingFind(tx, &lst, pi, order)
	return lst, po, err
}

func GetOperateLogList(q *OperateQ, pi *uv.PagingIn, order *uv.Order) ([]model.OperateLog, *uv.PagingOut, error) {
	tx := q.FilterQuery(db.DB)
	var lst = make([]model.OperateLog, 0)

	po, err := uv.PagingFind(tx, &lst, pi, order)
	return lst, po, err
}
