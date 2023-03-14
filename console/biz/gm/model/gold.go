package model

import (
	"console/mods/game_db"
)

type Gold struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code int    `json:"code"`
	Name string `json:"name" gorm:"type:string;size:128"`
}

func init() {
	game_db.RegTables(TaiwanCain2nd, &Gold{})
}
