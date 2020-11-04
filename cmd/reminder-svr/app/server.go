/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:04
 */
package app

import (
	reminder_apisvr "github.com/myxy99/reminder/internal/reminder-apisvr"
	"net/http"
)

func Run(stopCh <-chan struct{}) error {
	apiServer := NewApiServer()
	err := apiServer.PrepareRun(stopCh)
	if err != nil {
		return err
	}
	return apiServer.Run(stopCh)
}

func NewApiServer() *reminder_apisvr.WebServer {
	return &reminder_apisvr.WebServer{Server: new(http.Server)}
}
