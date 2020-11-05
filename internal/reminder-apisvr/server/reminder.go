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
	"github.com/myxy99/reminder/pkg/lunar"
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
	Time    int
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
	)
	var message = make(Message)
	var l = new(lunar.Lunar)

	timeRepository = impl.NewTimeRepository(r.db.DB())

	//todo 日期检索

	for k, t := range []int{1} {
		var (
			timer           time.Time
			month, day      int
			data, lunarData []*models.Time
		)
		timer = time.Now().AddDate(0, 0, t)
		month, _ = strconv.Atoi(fmt.Sprintf("%d", timer.Month()))
		day = timer.Day()
		//阳
		data, err = timeRepository.GetByTime(1, month, day, 1, k+1)
		fmt.Println(1, month, day, 1, k+1,"&&&&")
		//阴
		_, month, day = l.ToLunar(timer.Format("20060102"))
		fmt.Println(2, month, day, 1, k+1,"....")
		lunarData, err = timeRepository.GetByTime(2, month, day, 1, k+1)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, v := range append(data, lunarData...) {
			message[v.User.Email] = append(message[v.User.Email], &Data{
				Message: v.Message,
				Time:    t,
			})
		}
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
