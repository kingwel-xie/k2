package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/casbin"
	"github.com/kingwel-xie/k2/common/config"
	toolsDB "github.com/kingwel-xie/k2/core/tools/database"
	"github.com/kingwel-xie/k2/core/utils"
)

// Setup 配置数据库
func Setup() {
	setupSimpleDatabase(config.DatabaseConfig)
}

func setupSimpleDatabase(c *config.Database) {
	open := opens[c.Driver]
	if open == nil {
		log.Fatalf("invalid DB driver '%s'", c.Driver)
	}

	//log.Infof("db => %s", utils.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		CreateBatchSize: 500, // set to 500 to solve 'too many SQL variables'
		Logger:          &gormLogger{SlowThreshold: 200 * time.Millisecond},
	}, open)

	if err != nil {
		log.Fatal(utils.Red(c.Driver+" connect error :"), err)
	} else {
		log.Info(utils.Green(c.Driver + " connect success !"))
	}

	log.Infof("db connected, dialector = '%s'", db.Dialector.Name())

	common.Runtime.SetDb(db)

	// set up the casbin policy for Authenticator
	e := casbin.Setup(db)
	common.Runtime.SetCasbin(e)
}
