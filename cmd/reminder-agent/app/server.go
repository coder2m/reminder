/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:00
 */
package app

import (
	reminder_agent "github.com/myxy99/reminder/internal/reminder-agent"
)

func Run(stopCh <-chan struct{}) error {
	server := NewServer()
	err := server.PrepareRun(stopCh)
	if err != nil {
		return err
	}
	return server.Run(stopCh)
}

func NewServer() *reminder_agent.Server {
	return &reminder_agent.Server{}
}
