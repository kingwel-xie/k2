package migrate

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/database"
	"github.com/kingwel-xie/k2/common/models"
	"github.com/kingwel-xie/k2/core/migration"
	"github.com/kingwel-xie/k2/core/utils"
)

var (
	configYml string
	generate  bool
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

// fixme 在您看不见代码的时候运行迁移，我觉得是不安全的，所以编译后最好不要去执行迁移
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
}

func run() {
	if !generate {
		fmt.Println(`start init`)
		//1. 读取配置
		config.Setup(
			configYml,
			initDB,
		)
	} else {
		fmt.Println(`generate migration file`)
		_ = genFile()
	}
}

func migrateModel() error {
	db := common.Runtime.GetDb()
	if config.DatabaseConfig.Driver == "mysql" {
		//初始化数据库时候用
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	err := db.AutoMigrate(&models.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}
func initDB() {
	//3. 初始化数据库链接
	database.Setup()
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"

	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	utils.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")

	return nil
}
