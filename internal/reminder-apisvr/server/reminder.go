/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 16:22
 */
package server

import (
	"encoding/json"
	"fmt"
	"github.com/myxy99/reminder/internal/reminder-apisvr/models"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories/impl"
	"github.com/myxy99/reminder/pkg/client/database"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
	"log"
	"strconv"
	"time"
)

type reminder struct {
	db *database.Client
	Mq *rabbitmq.RabbitMQ
}

type Data struct {
	Message string
	TimeNum    int
}

type Message map[string][]*Data

func NewReminder(db *database.Client, Mq *rabbitmq.RabbitMQ) *reminder {
	return &reminder{db, Mq}
}

func (r *reminder) Run() {
	var (
		timeRepository repositories.TimeRepository
		err            error
		dataMsg        []byte
		timer          time.Time
		month, day     int
		data           []*models.Time
	)
	var message = make(Message)
	timeRepository = impl.NewTimeRepository(r.db.DB())
	//todo 日期检索
	timer = time.Now()
	month, _ = strconv.Atoi(fmt.Sprintf("%d", timer.Month()))
	day = timer.Day()
	data, err = timeRepository.GetByTime(month, day)
	for _, v := range data {
		message[v.User.Email] = append(message[v.User.Email], &Data{
			Message: v.Remind.Message,
			TimeNum:    v.TimeNum,
		})
	}
	if len(message) == 0 {
		return
	}
	dataMsg, err = json.Marshal(message)
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
