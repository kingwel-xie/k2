package email

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
)

// Setup 配置Email服务
func Setup() {
	email := config.EmailConfig.Setup()
	common.Runtime.SetEmail(email)
}
