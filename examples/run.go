// +build examples

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kingwel-xie/k2/common"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp/inmg?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	syncEnforce := casbin.Setup(db)
	common.Runtime.SetDb(db)
	common.Runtime.SetCasbin(syncEnforce)

	e := gin.Default()
	common.Runtime.SetEngine(e)
	log.Fatal(e.Run(":8000"))
}
