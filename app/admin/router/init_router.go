package router

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/middleware"
	log "github.com/kingwel-xie/k2/core/logger"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit("")
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	// 注册业务路由
	InitServiceRouter(r, authMiddleware)
}
