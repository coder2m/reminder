/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 15:27
 */
package reminder

import (
	"github.com/robfig/cron/v3"
)

type CronServer struct {
	o *Options
	c *cron.Cron
}

func (r *CronServer) Run(stopCh <-chan struct{}, f cron.Job) {
	for _, v := range r.o.Spec {
		_, _ = r.c.AddJob(v, f)
	}
	r.c.Start()
	<-stopCh
	r.c.Stop()
}

func NewReminderClient(o *Options) *CronServer {
	return &CronServer{c: cron.New(), o: o}
}
