package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"

	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("cronjob")

type CronJob struct {
	cron    *cron.Cron
	jobList map[string]*Job
	lock    sync.Mutex
}

func Setup() *CronJob {
	return &CronJob{
		cron:    newWithSeconds(),
		jobList: make(map[string]*Job, 0),
	}
}

// newWithSeconds returns a Cron with the seconds field enabled.
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func (j *CronJob) Start() {
	j.cron.Start()
}

func (j *CronJob) AddJob(name string, expression string, exec func()) error {
	_, b := j.jobList[name]
	if b {
		return fmt.Errorf("job %s already existed", name)
	}

	cmd := &Job{
		Name:       name,
		Expression: expression,
		exec:       exec,
	}

	id, err := j.cron.AddJob(expression, cmd)
	if err != nil {
		return err
	}
	cmd.entryId = int(id)

	j.lock.Lock()
	j.jobList[name] = cmd
	j.lock.Unlock()
	return nil
}

func (j *CronJob) RemoveJob(name string) error {
	cmd, b := j.jobList[name]
	if !b {
		return fmt.Errorf("job %s not exists", name)
	}
	j.cron.Remove(cron.EntryID(cmd.entryId))
	return nil
}

// ListJobs returns a slice of all running jobs
func (j *CronJob) ListJobs() []*Job {
	s := make([]*Job, 0, len(j.jobList))
	for _, v := range j.jobList {
		s = append(s, v)
	}
	return s
}
