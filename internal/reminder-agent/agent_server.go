/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:01
 */
package reminder_agent

import (
	"context"
	"github.com/myxy99/reminder/internal/reminder-agent/config"
	"github.com/myxy99/reminder/internal/reminder-agent/server"
	"github.com/myxy99/reminder/pkg/client/email"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
	"log"
)

type Server struct {
	Config *config.Cfg
	Mq     *rabbitmq.RabbitMQ
	Email  *email.Email
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	err = s.installCfg()
	if err != nil {
		return
	}

	err = s.installRabbitMQ(stopCh)
	if err != nil {
		return
	}

	err = s.installEmail()
	if err != nil {
		return
	}

	return nil
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("running ~~ ")
	err = s.Mq.ConsumeSimple(server.NewServer(s.Email).Send, stopCh)

	return err
}

func (s *Server) installRabbitMQ(stopCh <-chan struct{}) (err error) {
	s.Mq, err = rabbitmq.NewRabbitMQSimple("reminder", s.Config.RabbitMq)
	go func() {
		<-stopCh
		s.Mq.Destory()
	}()
	return
}

func (s *Server) installCfg() (err error) {
	s.Config, err = config.TryLoadFromDisk()
	return
}

func (s *Server) installEmail() (err error) {
	s.Email = email.NewEmail(s.Config.Email)
	return
}
