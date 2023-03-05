package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"admin/apis"
	"github.com/kingwel-xie/k2/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

// 需认证的路由代码
func registerSysRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysRole{}
	r := v1.Group("/role").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/role-list", api.GetAll)
		r1.PUT("/role-status", api.Update2Status)
		r1.PUT("/role-datascope", api.Update2DataScope)
	}
}
