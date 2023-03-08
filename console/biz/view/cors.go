package view

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setCORS(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	r.Use(cors.New(config))
}
