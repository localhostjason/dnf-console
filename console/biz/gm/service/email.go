package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"fmt"
	"github.com/localhostjason/webserver/server/util/uv"
	"time"
)

// SendEmail 插入 2条记录 letter postal
func SendEmail(characNo int, email *Email) error {
	dbx := game_db.DBPools.Get(model.TaiwanCain2nd)

	err := dbx.Table("letter").Create(map[string]interface{}{
		"charac_no":        characNo,
		"send_charac_no":   0,
		"send_charac_name": "GM",
		"letter_text":      "Thanks!",
		"stat":             1,
	}).Error

	if err != nil {
		return err
	}

	type _R struct {
		LetterId int `json:"letter_id"`
	}
	var data _R
	dbx.Table("letter").Where("charac_no = ?", characNo).Order("letter_id desc").Take(&data)

	fmt.Println(123, data.LetterId)

	now := time.Now()
	return dbx.Debug().Table("postal").Create(map[string]interface{}{
		"occ_time":          now.Format("2006-01-02 15:04:05"),
		"send_charac_no":    0,
		"send_charac_name":  "GAME MASTER",
		"receive_charac_no": characNo,
		"item_id":           email.Code,
		"add_info":          email.Number,
		"letter_id":         data.LetterId,
	}).Error

}

func GetGoldList(q *GoldQ, pi *uv.PagingIn, order *uv.Order) ([]model.Gold, *uv.PagingOut, error) {
	dbx := game_db.DBPools.Get(model.TaiwanCain2nd)
	tx := q.FilterQuery(dbx)

	var lst = make([]model.Gold, 0)
	po, err := uv.PagingFind(tx, &lst, pi, order)

	return lst, po, err
}
