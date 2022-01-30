package sms

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
)


// Setup 配置SMS服务
func Setup() {
	sms := config.SmsConfig.Setup()
	common.Runtime.SetSms(sms)
}
