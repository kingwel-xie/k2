package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysInboxRouter)
}

// registerSysInboxRouter
func registerSysInboxRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysInbox{}
	r := v1.Group("/inbox").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		//r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/unread", api.GetUnread)
		r.POST("/read", api.MarkRead)
	}
}