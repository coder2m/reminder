/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 16:22
 */
package server

import (
	"encoding/json"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories/impl"
	"github.com/myxy99/reminder/pkg/client/database"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
	"log"
)

type reminder struct {
	db *database.Client
	Mq *rabbitmq.RabbitMQ
}

type Message struct {
	To  string
	msg string
}

func NewReminder(db *database.Client, Mq *rabbitmq.RabbitMQ) *reminder {
	return &reminder{db, Mq}
}

func (r *reminder) Run() {
	var timeRepository repositories.TimeRepository
	timeRepository = impl.NewTimeRepository(r.db.DB())
	data, err := timeRepository.GetByTime(1, 11, 5)

	if err != nil {
		log.Println(err.Error())
		return
	}

	var message Message

	for _, v := range data {
		message = Message{
			To:  v.User.Email,
			msg: v.Message,
		}

		dataMsg, err := json.Marshal(message)

		if err != nil {
			log.Println(err.Error())
			return
		}

		err = r.Mq.PublishSimple(dataMsg)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}

}
