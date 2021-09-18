package admin

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/kingwel-xie/k2/app/admin/models"
	"github.com/kingwel-xie/k2/app/admin/router"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/database"
	"github.com/kingwel-xie/k2/common/global"
	"github.com/kingwel-xie/k2/common/middleware"
	"github.com/kingwel-xie/k2/common/storage"
	log "github.com/kingwel-xie/k2/core/logger"
	"github.com/kingwel-xie/k2/core/utils"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "admin",
		Short:        "Start API server",
		Example:      "github.com/kingwel-xie/k2 admin -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

func setup() {
	//1. 读取配置
	config.Setup(
		configYml,
		database.Setup,
		storage.Setup,
	)
	//注册监听函数
	queue := common.Runtime.Queue()
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperaLog)
	queue.Register(global.ApiCheck, models.SaveSysApi)

	// DEV tool for convenience, we can quickly add/modify database...
	initDB()

	log.Info(`starting api server...`)
}

func initDB() {
	log.Info(`migrating database...`)

	db := common.Runtime.GetDb()

	// run for db migration
	_ = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(
			new(models.TbxCountry),
			//new(models.TbxAddress),
		)
		if err != nil {
			log.Fatalf("cannot migrate DB, %v", err)
		}
		return nil
	})
}

func run() error {
	if config.ApplicationConfig.Mode == utils.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: common.Runtime.GetEngine(),
	}

	if apiCheck {
		var routers = common.Runtime.GetRouter()
		mp := make(map[string]interface{}, 0)
		mp["List"] = routers
		message, err := common.Runtime.GetStreamMessage("", global.ApiCheck, mp)
		if err != nil {
			log.Infof("GetStreamMessage error, %s \n", err.Error())
			//日志报错错误，不中断请求
		} else {
			err = common.Runtime.Queue().Append(message)
			if err != nil {
				log.Infof("Append message error, %s \n", err.Error())
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接
		if config.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	fmt.Println(utils.Red(string(global.LogoContent)))
	tip()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Println(utils.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/index.html \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/swagger/index.html \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", utils.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", utils.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")

	return nil
}

func tip() {
	usageStr := `欢迎使用 ` + utils.Green(`k2 `+global.Version) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		common.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
	}
	if config.SslConfig.Enable {
		r.Use(middleware.TlsHandler())
	}
	r.Use(middleware.Sentinel())
	r.Use(middleware.RequestId(utils.TrafficKey))

	middleware.InitMiddleware(r)
}
