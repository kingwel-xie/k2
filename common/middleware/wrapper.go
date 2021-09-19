package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinWrapper http.Handler 转换成 gin.HandlerFunc
func GinWrapper(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
