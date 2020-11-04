/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:14
 */
package reminder_apisvr

import (
	"context"
	v1 "github.com/myxy99/reminder/internal/reminder-apisvr/api/v1"
	"github.com/myxy99/reminder/internal/reminder-apisvr/config"
	myValidator "github.com/myxy99/reminder/internal/reminder-apisvr/validator"
	"github.com/myxy99/reminder/pkg/client/database"
	"github.com/myxy99/reminder/pkg/validator"
	"log"
	"net/http"
)

type WebServer struct {
	DB *database.Client

	Config *config.Cfg

	Server *http.Server

	Validator *validator.Validator
}

func (s *WebServer) PrepareRun(stopCh <-chan struct{}) (err error) {
	err = s.installCfg()
	if err != nil {
		return
	}

	err = s.installDatabase(stopCh)
	if err != nil {
		return
	}

	s.installHttpServer()

	err = s.installValidator()
	if err != nil {
		return
	}

	s.installAPIs()
	return nil
}

func (s *WebServer) Run(stopCh <-chan struct{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
		_ = s.Server.Shutdown(ctx)
	}()
	log.Printf("Start listening on %s", s.Server.Addr)
	err = s.Server.ListenAndServe()
	return err
}

func (s *WebServer) migration() {
	s.DB.DB().AutoMigrate(
		//new(models.Auth),
	)
}

func (s *WebServer) installAPIs() {
	s.Server.Handler = v1.InitRouter(s.DB, s.Validator)
}

func (s *WebServer) installHttpServer() {
	s.Server.Addr = s.Config.Server.Addr
}

func (s *WebServer) installValidator() error {
	s.Validator = validator.New()
	return s.Validator.InitTrans(s.Config.Server.Locale, myValidator.RegisterValidation)
}

func (s *WebServer) installDatabase(stopCh <-chan struct{}) (err error) {
	s.DB, err = database.NewDatabaseClient(s.Config.Database, stopCh)
	s.migration()
	return
}

func (s *WebServer) installCfg() (err error) {
	s.Config, err = config.TryLoadFromDisk()
	return
}
