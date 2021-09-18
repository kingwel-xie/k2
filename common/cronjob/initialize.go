package cronjob

import (
	"log"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/core/cronjob"
)

// 初始化
func Setup() {
	log.Println("Job Starting...")
	crontab := cronjob.Setup()

	common.Runtime.SetCrontab(crontab)

	crontab.Start()
	log.Println("Job started successfully.")
}

//
//// ListJobs returns a slice of all running jobs
//func ListJobs() []*Job {
//	s := make([]*Job, 0, len(jobList))
//	for _, v := range jobList {
//		s = append(s, v)
//	}
//	return s
//}