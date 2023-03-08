package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/ue"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"sync"
)

var lock sync.Mutex

// ErrorHandler remove
func ErrorHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*ue.Error); ok {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
			} else {
				log.Error(string(debug.Stack()))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}

	}()
	lock.Lock()
	defer lock.Unlock()
	c.Next()
}
