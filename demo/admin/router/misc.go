package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerNoAuthRouter)
	routerCheckRole = append(routerCheckRole, registerAuthRouter)
}

func registerAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	miscApi := apis.TbxMisc{}

	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/all-dict", miscApi.GetAll)
	}
}

func registerNoAuthRouter(v1 *gin.RouterGroup) {
	miscApi := apis.TbxMisc{}

	v1.GET("/limited-download/:uuid", miscApi.LimitedDownload)

	// check username existence
	v1.GET("/check-existence/:username", apis.SysUser{}.CheckExistence)

	v1.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler()
	})
}
