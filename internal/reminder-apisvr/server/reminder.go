/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 16:22
 */
package server

import (
	"fmt"
	"github.com/myxy99/reminder/pkg/client/database"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
)

type reminder struct {
	db *database.Client
	Mq *rabbitmq.RabbitMQ
}

func NewReminder(db *database.Client, Mq *rabbitmq.RabbitMQ) *reminder {
	return &reminder{db, Mq}
}

func (r *reminder) Run() {
	//var timeRepository repositories.TimeRepository
	//timeRepository = impl.NewTimeRepository(r.db.DB())
	fmt.Println("Reminder running!!!!")
}
