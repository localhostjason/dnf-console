package model

import "console/mods/game_db"

type Accounts struct {
	Uid         int    `json:"uid" gorm:"column:UID;primaryKey"`
	AccountName string `json:"account_name" gorm:"column:accountname;type:string;size:128;unique;not null"`
	Password    string `json:"-" gorm:"type:string;size:128;not null"`
	QQ          string `json:"qq" gorm:"type:string;size:128;"`
}

func init() {
	game_db.RegTables(DTaiwan, &Accounts{})
}
