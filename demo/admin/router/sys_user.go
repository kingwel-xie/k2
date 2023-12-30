package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"admin/apis"
	"github.com/kingwel-xie/k2/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/pwd/reset", api.ResetPwd)
		r.PUT("/token/reset", api.ResetApiToken)
		r.PUT("/status", api.UpdateStatus)
	}

	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/list", api.ListNoCheck)
		user.GET("/profile", api.GetProfile)
		user.POST("/profile", api.UpdateProfile)
		user.POST("/avatar", api.InsetAvatar)
		user.PUT("/pwd/set", api.UpdatePwd)
		user.PUT("/token", api.RegenerateApiToken)
	}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", api.GetInfo)
	}
}
