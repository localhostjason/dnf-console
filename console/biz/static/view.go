package static

import (
	"console/mods/pathx"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AddStaticToRouter(r *gin.Engine) {
	exeDir, _ := pathx.GetExeDir()

	r.GET("/", redirectRoot)
	r.Static("/static", filepath.Join(exeDir, "web", "static"))
}

func redirectRoot(c *gin.Context) {
	c.Redirect(302, "/static/index.html")
}
