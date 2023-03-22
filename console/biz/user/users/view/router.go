package view

import (
	"console/mods/ginx"
	"console/mods/pathx"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"path/filepath"
)

func InitUsersRouter(r *ginx.RouterGroup) {
	userInfo := r.Group("用户个人信息", "info")
	{
		userInfo.GET("获取个人信息", "", getUserInfo)
		userInfo.PUT("更新个人信息", "", updateUserInfo)
	}

	users := r.Group("用户管理", "list")
	{
		users.GET("获取用户列表", "", getUsers)

	}

	r.GET("下载", "file", loadFile)
}

func loadFile(c *gin.Context) {
	//uv.PEIf(E_USER_INFO_UPDATE, errors.New("test error"))
	pwd, _ := pathx.GetExeDir()
	file := filepath.Join(pwd, "config", "rbac_model.conf")
	output := url.QueryEscape(filepath.Base(file))
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", output))
	c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(file)

}
