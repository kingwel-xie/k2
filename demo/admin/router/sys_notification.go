package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common/middleware"

	"admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysNotificationRouter)
}

// registerSysNotificationRouter
func registerSysNotificationRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysNotification{}
	r := v1.Group("/notification").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		//r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}