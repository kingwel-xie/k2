package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysOutboxRouter)
}

// registerSysOutboxRouter
func registerSysOutboxRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysOutbox{}
	r := v1.Group("/outbox").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
	}
}