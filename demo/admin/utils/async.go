package utils

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/core/logger"
	"gorm.io/gorm"
)

var asyncTaskLog = logger.Logger("async-task")

func AsyncTask(name string, task func(log logger.StandardLogger) error) {
	go func() {
		log := asyncTaskLog.WithFields("task", name)
		if err := task(log); err != nil {
			asyncTaskLog.Warnf("failed to execute async task [%s]: %s", name, err.Error())
		}
	}()
}

func AsyncDBTask(name string, task func(log logger.StandardLogger, db *gorm.DB) error) {
	AsyncTask(name, func(log logger.StandardLogger) error {
		return task(log, common.Runtime.GetDb())
	})
}
