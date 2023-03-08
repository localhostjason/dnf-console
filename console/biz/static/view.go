package static

import (
	"github.com/localhostjason/webserver/server/config"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AddStaticToRouter(r *gin.Engine) {
	r.GET("/", redirectRoot)
	r.Static("/static", filepath.Join(config.RootDir, "web", "static"))
}

func redirectRoot(c *gin.Context) {
	c.Redirect(302, "/static/index.html")
}
