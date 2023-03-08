package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
	"net/http"
	"path"
)

type RouterGroup struct {
	RouterGroup *gin.RouterGroup
	Name        string
}

func NewRouterGroup(r *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{RouterGroup: r}
}

func (r *RouterGroup) Group(name string, relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{
		RouterGroup: r.RouterGroup.Group(relativePath, handlers...),
		Name:        name,
	}
}

// POST is a shortcut for router.Handle("POST", path, handle).
func (r *RouterGroup) POST(name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	r.authzHandler(name, r.Name, relativePath, http.MethodPost, r.RouterGroup)
	return r.RouterGroup.POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (r *RouterGroup) GET(name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	r.authzHandler(name, r.Name, relativePath, http.MethodGet, r.RouterGroup)
	return r.RouterGroup.GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (r *RouterGroup) DELETE(name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	r.authzHandler(name, r.Name, relativePath, http.MethodDelete, r.RouterGroup)
	return r.RouterGroup.DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (r *RouterGroup) PATCH(name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	r.authzHandler(name, r.Name, relativePath, http.MethodPatch, r.RouterGroup)
	return r.RouterGroup.PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (r *RouterGroup) PUT(name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	r.authzHandler(name, r.Name, relativePath, http.MethodPut, r.RouterGroup)
	return r.RouterGroup.PUT(relativePath, handlers...)
}

func (r *RouterGroup) authzHandler(name string, groupName string, relativePath string, method string, route *gin.RouterGroup) {
	if groupName == "" || method == "" {
		return
	}
	var authz []Authz
	url := path.Join(route.BasePath(), relativePath)
	z := Authz{GroupName: groupName, ApiName: name, Url: url, Method: method}
	db.DB.Limit(1).Find(&authz, z)

	if len(authz) == 0 {
		db.DB.Create(&z)
	}
}
