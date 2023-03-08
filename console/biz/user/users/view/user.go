package view

import (
	"console/biz/user/users/model"
	"console/biz/user/users/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getUserInfo(c *gin.Context) {
	c.Set(operateKey, uv.OP(I_OP, "1", "hello"))

	data := service.GetUserInfo(c)
	c.JSON(200, data)
}

func updateUserInfo(c *gin.Context) {
	userQ := &model.UserInfoPut{}
	uv.PB(c, userQ)

	err := service.UpdateUserInfo(c, userQ.Desc)
	uv.PEIf(E_USER_INFO_UPDATE, err)
	c.Status(201)
}

func getUsers(c *gin.Context) {
	data := service.GetUsers()
	c.JSON(200, data)
}
