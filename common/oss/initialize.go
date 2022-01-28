package oss

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
)


// Setup 配置OSS服务
func Setup() {
	oss := config.OssConfig.Setup()
	common.Runtime.SetOss(oss)
}
