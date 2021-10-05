package router

import (
	"mime"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"admin/apis"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/middleware"
	"github.com/kingwel-xie/k2/core/ws"

	_ "admin/docs"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}
	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)

	r.NoRoute(func(context *gin.Context) {
		context.Redirect(302, "/site/")
	})
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()
	if config.ApplicationConfig.Mode != "prod" {
		//r.GET("/", func(c *gin.Context) {
		//	c.Redirect(302, "/site/")
		//	//c.Status(http.StatusOK)
		//})
	}
	r.GET("/info", middleware.Ping)

	r.GET("/metrics", middleware.GinWrapper(promhttp.Handler()))
	//健康检查
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	r.Static("/site", "./site")
	r.Static("/static", "./static")
	if config.ApplicationConfig.Mode != "prod" {
		r.Static("/form-generator", "./static/form-generator")
	}
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}
	info := r.Group("/server-monitor").Use(authMiddleware.MiddlewareFunc())
	{
		info.GET("", middleware.ServerInfo)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)

		v1.GET("/captcha", middleware.GenerateCaptchaHandler)
	}

	registerBaseRouter(v1, authMiddleware)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}
	api2 := apis.SysDept{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/roleMenuTreeSelect/:roleId", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeSelect/:roleId", api2.GetDeptTreeRoleSelect)

		v1auth.POST("/logout", authMiddleware.LogoutHandler)
		var api = middleware.File{}
		v1auth.POST("/public/uploadFile", api.UploadFile)
	}
}
