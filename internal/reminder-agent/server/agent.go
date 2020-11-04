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

type Message struct {
	To  string
	msg string
}

func NewServer(email *email.Email) *Server {
	return &Server{email}
}

func (s *Server) Send(data []byte) (err error) {
	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		return
	}
	err = s.emailClient.SendEmail([]string{message.To}, "提示", message.msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return err
}
