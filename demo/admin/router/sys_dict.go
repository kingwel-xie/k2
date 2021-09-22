package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"admin/apis"
	"github.com/kingwel-xie/k2/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDictRouter)
}

func registerDictRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	dictApi := apis.SysDictType{}
	dataApi := apis.SysDictData{}
	dicts := v1.Group("/dict").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		dicts.GET("/data", dataApi.GetPage)
		dicts.GET("/data/:dictCode", dataApi.Get)
		dicts.POST("/data", dataApi.Insert)
		dicts.PUT("/data/:dictCode", dataApi.Update)
		dicts.DELETE("/data", dataApi.Delete)

		dicts.GET("/type", dictApi.GetPage)
		dicts.GET("/type/:id", dictApi.Get)
		dicts.POST("/type", dictApi.Insert)
		dicts.PUT("/type/:id", dictApi.Update)
		dicts.DELETE("/type", dictApi.Delete)
	}
	typeSelect := v1.Group("/dict").Use(authMiddleware.MiddlewareFunc())
	{
		typeSelect.GET("/type-option-select", dictApi.GetAll)
	}
	opSelect := v1.Group("/dict-data").Use(authMiddleware.MiddlewareFunc())
	{
		opSelect.GET("/option-select", dataApi.GetAll)
	}
}
