package service

import (
	roleService "console/biz/user/role/service"
	"console/biz/user/users/model"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
)

func GetUserInfo(c *gin.Context) *model.User {
	return roleService.GetCurrentUser(c)
}

func UpdateUserInfo(c *gin.Context, desc string) error {
	currentUser := roleService.GetCurrentUser(c)
	currentUser.Desc = desc

	return db.DB.Save(currentUser).Error
}

func GetUsers() []model.User {
	var users = make([]model.User, 0)
	db.DB.Find(&users)
	return users
}
