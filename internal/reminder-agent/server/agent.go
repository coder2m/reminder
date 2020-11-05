/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:15
 */
package server

import (
	"encoding/json"
	"github.com/myxy99/reminder/pkg/client/email"
)

type Server struct {
	emailClient *email.Email
}

type Data struct {
	Message string
	Time    int
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
			msgStr += "<li>" + s.typeToMsg(msg.Time) + msg.Message + "</li><br>"
		}
		err = s.emailClient.SendEmail([]string{k}, "reminder", msgStr)
	}

	return err
}

func (s *Server) typeToMsg(timeType int) string {
	switch timeType {
	case 1:
		return "一天后："
	case 3:
		return "三天后："
	case 5:
		return "五天后："
	default:
		return ""
	}
}
