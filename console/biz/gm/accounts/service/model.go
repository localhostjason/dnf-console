package service

import "console/biz/gm/accounts/model"

type AccountResult struct {
	model.Accounts
	Roles int64 `json:"roles"`
	Money any   `json:"money"`
}
