package model

import "time"

type CharacInfo struct {
	MId        int       `json:"m_id"`
	CharacNo   int       `json:"charac_no"`
	CharacName string    `json:"charac_name"`
	Lev        int       `json:"lev"`
	CreateTime time.Time `json:"create_time"`
}
