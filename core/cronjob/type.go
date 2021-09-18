package cronjob

import (
	"time"
)

type JobExec interface {
	Exec(arg interface{}) error
}

type Job struct {
	Name       string
	Expression string
	entryId    int
	exec       func()
}

func (j *Job) Run() {
	startTime := time.Now()

	j.exec()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	log.Debugf("run job %s, took %v", j.Name, latencyTime)
}
