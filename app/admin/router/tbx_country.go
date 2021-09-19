package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/app/admin/apis"
	"github.com/kingwel-xie/k2/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTbxCountryRouter)
}

// registerTbxCountryRouter
func registerTbxCountryRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.TbxCountry{}
	r := v1.Group("/country").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:code", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:code", api.Update)
		r.DELETE("", api.Delete)
	}
}
