/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:15
 */
package server

import (
	"encoding/json"
	"fmt"
	"github.com/myxy99/reminder/pkg/client/email"
)

type Server struct {
	emailClient *email.Email
}

type Data struct {
	Message string
	TimeNum int
}

type Message map[string][]*Data

func NewServer(email *email.Email) *Server {
	return &Server{email}
}

func (s *Server) Send(data []byte) (err error) {
	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		return
	}
	for k, v := range message {
		var msgStr string
		for _, msg := range v {
			msgStr += fmt.Sprintf("<li>%d天后:%s</li><br>", msg.TimeNum, msg.Message)
		}
		err = s.emailClient.SendEmail([]string{k}, "reminder", msgStr)
	}

	return err
}
