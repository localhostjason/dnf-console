package service

import (
	"console/biz/user/auth/model"
	m "console/biz/user/users/model"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const currentUserKey = "current_user"
const currentPassword = "current_password"
const loginFailedKey = "___login_failed"

var _c model.ConfigAuth

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*m.User); ok {
		return jwt.MapClaims{
			_c.IDKey: v.JwtKey,
		}
	}
	return jwt.MapClaims{}
}

func IdHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	jwtKey := claims[_c.IDKey].(string)

	var user = &m.User{}
	err := db.DB.Where("jwt_key = ?", jwtKey).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

type loginArgs struct {
	Username string `json:"username" binding:"required,alphanum,lte=32"`
	Password string `json:"password" binding:"required,printascii,lte=128"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginValues loginArgs
	if err := c.ShouldBind(&loginValues); err != nil {
		return nil, errors.New("请正确输入用户名密码")
	}

	userName := loginValues.Username
	password := loginValues.Password

	// 直接记录下来， 不管成功与否， 后面看情况使用
	c.Set(loginFailedKey, &m.User{Username: userName})

	// 密码登录
	var user m.User
	err := db.DB.Where("username = ?", userName).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || !user.CheckPassword(password) {
		return nil, errors.New("用户名或者密码填写不对")
	}

	c.Set(currentUserKey, &user)
	c.Set(currentPassword, password)
	return &user, nil

}

func Authorizator(data interface{}, c *gin.Context) bool {
	if data == nil {
		return false
	}
	fmt.Println("data", data)
	return true
}

// UnAuth 密码登录失败时候调用的函数
func UnAuth(c *gin.Context, code int, message string) {
	desc := "登录失败"
	msg := message
	if message == "Token is expired" {
		desc = ""
		msg = "您的登录已过期，请重新登录"
	}

	if message == "query token is empty" {
		desc = ""
		msg = "未携带身份凭证"
	}

	if message == "you don't have permission to access this resource" {
		msg = "携带身份凭证不正确，认证失败"
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"msg":  msg,
		"desc": desc,
	})
}

// LoginResponse 密码登录成功时调用的函数
func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	u, _ := c.Get(currentUserKey)
	user := u.(*m.User)

	info := user.GetInfo()

	info["token"] = token

	now := time.Now()
	user.LastLoginTime = &now
	db.DB.Save(user)

	c.JSON(http.StatusOK, info)
}

// LogoutResponse 退出登录
func LogoutResponse(c *gin.Context, code int) {
	c.Status(201)
}

func CurrentUser(c *gin.Context) *m.User {
	u := IdHandler(c)
	currentUser := u.(*m.User)
	return currentUser
}
