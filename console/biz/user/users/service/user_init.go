package service

import (
	"console/biz/user/users/model"
	"github.com/localhostjason/webserver/db"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	Normal = 1
	Locked = -1
)

const (
	ADMIN    = "admin"
	SECURITY = "security"
	AUDITOR  = "auditor"
)

const _defaultPassword = "123"

func InitUser() error {
	now := time.Now()
	users := []model.User{
		{
			Username: "admin",
			Role:     ADMIN,
			Email:    "",
			Desc:     "系统管理员",
			Time:     now,
		},
	}

	for i := range users {
		u := &users[i]

		var userList []model.User
		if db.DB.Limit(1).Where("username = ? ", u.Username).Find(&userList); len(userList) == 0 {
			u.SetPassword(_defaultPassword)
			u.JwtKey = uuid.NewV4()
			db.DB.Create(u)
		}
	}
	return nil
}

func init() {
	db.AddInitHook(InitUser)
}
