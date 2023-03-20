package service

import (
	logModel "console/biz/log/model"
	"fmt"
	"github.com/localhostjason/webserver/db"
	"time"
)

func GetDashChart() ChartResult {
	return ChartResult{
		Cera:      getChartByType("cera"),
		CeraPoint: getChartByType("cera_point"),
	}
}

func getChartByType(typ string) ChartTranInfo {
	var data []ChartInfo

	now := time.Now()
	nowYear := now.Format("2006")

	tx := db.DB.Debug().Model(&logModel.RechargeLog{}).
		Select("sum(number) as total,strftime('%Y',time) as year, strftime('%m',time) AS month").
		Group("year, month").
		Having("year = ?", nowYear)

	if typ == "cera" {
		tx = tx.Where("action = ?", logModel.RechargeCera)
	} else {
		tx = tx.Where("action = ?", logModel.RechargeCeraPoint)
	}
	tx.Find(&data)

	date := make([]string, 0)
	for i := 1; i <= 12; i++ {
		date = append(date, fmt.Sprintf("%d月", i))
	}

	number := make([]int, 0)
	for i := 1; i <= 12; i++ {
		isAppend := false
		for _, info := range data {
			if info.Month == i {
				number = append(number, info.Total)
				isAppend = true
				break
			}
		}

		if !isAppend {
			number = append(number, 0)
		}

	}

	return ChartTranInfo{
		Date:  date,
		Total: number,
	}
}
