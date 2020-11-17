/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:04
 */
package app

import (
	"fmt"
	reminder_apisvr "github.com/myxy99/reminder/internal/reminder-apisvr"
	"github.com/myxy99/reminder/pkg/log"
	"net/http"
)

func Run(stopCh <-chan struct{}) error {
	apiServer := NewApiServer()
	err := apiServer.PrepareRun(stopCh)
	if err != nil {
		log.Info(fmt.Sprintf("apiServer Start PrepareRun err %s", err.Error()))
		return err
	}
	return apiServer.Run(stopCh)
}

func NewApiServer() *reminder_apisvr.WebServer {
	return &reminder_apisvr.WebServer{Server: new(http.Server)}
}
