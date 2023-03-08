package view

import (
	"console/biz/user/auth/model"
	"console/biz/user/auth/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

func AddJwtAuth(authApi, authedApi *gin.RouterGroup) error {
	m, err := newAuthMiddleWare()
	if err != nil {
		return err
	}

	authApi.POST("login", m.LoginHandler)
	authedApi.Use(m.MiddlewareFunc())

	authedApi.POST("auth/logout", m.LogoutHandler)
	return nil
}

func newAuthMiddleWare() (*jwt.GinJWTMiddleware, error) {
	conf, err := model.GetConfig()
	if err != nil {
		return nil, err
	}

	c := authConfig(conf)
	return jwt.New(c)
}

func authConfig(conf model.ConfigAuth) *jwt.GinJWTMiddleware {
	c := jwt.GinJWTMiddleware{
		Realm:           conf.Realm,
		Key:             []byte(conf.Secret),
		Timeout:         time.Duration(conf.Timeout) * time.Second,
		MaxRefresh:      time.Duration(conf.MaxRefresh) * time.Second,
		IdentityKey:     conf.IDKey,
		PayloadFunc:     service.PayloadFunc,
		IdentityHandler: service.IdHandler,
		Authenticator:   service.Authenticator,
		Authorizator:    service.Authorizator,
		Unauthorized:    service.UnAuth,
		LoginResponse:   service.LoginResponse,
		LogoutResponse:  service.LogoutResponse,
		TokenLookup:     "header: Authorization, query: token",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	}
	return &c
}
