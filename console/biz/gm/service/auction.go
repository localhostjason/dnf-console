package service

import (
	"console/biz/gm/model"
	"console/mods/game_db"
	"errors"
	"fmt"
	"time"
)

/*
1. 拍卖行 开启
2. 金币寄售 开启
3. todo 定时
*/

type StateResp struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Visible bool   `json:"visible"`
}

func GetAuctionState() []StateResp {

	var data = make([]StateResp, 0)

	data = append(data, StateResp{
		Name:    "gold",
		Desc:    "拍卖行",
		Visible: getGoldState(),
	})

	data = append(data, StateResp{
		Name:    "cera",
		Desc:    "金币寄售",
		Visible: getCeraState(),
	})
	return data
}

func getGoldState() bool {
	var visible = true
	now := time.Now()
	nowStr := now.Format("200601")
	dbx := game_db.DBPools.Get(model.TaiwanCainAuctionGold)

	tables := []string{"auction_history", "auction_history_buyer"}

	for _, table := range tables {
		t := fmt.Sprintf("%s_%s", table, nowStr)
		fmt.Println("table:", t)
		ok := dbx.Migrator().HasTable(t)
		if !ok {
			visible = false
			break
		}
	}
	return visible
}

func getCeraState() bool {
	var visible = true
	now := time.Now()
	nowStr := now.Format("200601")
	dbx := game_db.DBPools.Get(model.TaiwanCainAuctionCera)

	tables := []string{"auction_history", "auction_history_buyer"}

	for _, table := range tables {
		t := fmt.Sprintf("%s_%s", table, nowStr)
		ok := dbx.Migrator().HasTable(t)
		if !ok {
			visible = false
			break
		}
	}
	return visible
}

func OpenAuction(name string) error {
	if name != "gold" && name != "cera" {
		return errors.New("参数不对")
	}

	// 拍卖行
	if name == "gold" {
		visible := getGoldState()
		if visible {
			return nil
		}
		return openGold()
	}

	// 金币
	visible := getCeraState()
	if visible {
		return nil
	}
	return openCera()
}

func openGold() error {
	now := time.Now()
	nowStr := now.Format("200601")
	tables := []string{"auction_history", "auction_history_buyer"}
	dbx := game_db.DBPools.Get(model.TaiwanCainAuctionGold)

	for _, table := range tables {
		t := fmt.Sprintf("%s_%s", table, nowStr)
		ok := dbx.Migrator().HasTable(t)
		if ok {
			continue
		}

		if err := dbx.Debug().Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (LIKE %s);", t, table)).Error; err != nil {
			return err
		}
	}
	return nil
}

func openCera() error {
	now := time.Now()
	nowStr := now.Format("200601")
	tables := []string{"auction_history", "auction_history_buyer"}
	dbx := game_db.DBPools.Get(model.TaiwanCainAuctionCera)

	for _, table := range tables {
		t := fmt.Sprintf("%s_%s", table, nowStr)
		ok := dbx.Migrator().HasTable(t)
		if ok {
			continue
		}

		if err := dbx.Debug().Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (LIKE %s);", t, table)).Error; err != nil {
			return err
		}
	}
	return nil
}
