package app

import (
	"context"
	"github.com/BurntSushi/toml"
	xapp "github.com/coder2m/component"
	"github.com/coder2m/component/pkg/xconsole"
	"github.com/coder2m/component/pkg/xdefer"
	"github.com/coder2m/component/pkg/xflag"
	"github.com/coder2m/component/pkg/xvalidator"
	"github.com/coder2m/component/xcfg"
	"github.com/coder2m/component/xcfg/datasource/manager"
	"github.com/coder2m/component/xgovern"
	"github.com/coder2m/component/xinvoker"
	"github.com/coder2m/component/xmonitor"
	"github.com/coder2m/component/xtrace"
	"github.com/coder2m/reminder/internal/app/api/v1/registry"
	myValidator "github.com/coder2m/reminder/internal/app/validator"
	"net/http"
	"sync"
)

type Server struct {
	Server *http.Server
	err    error
	*sync.WaitGroup
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	s.initCfg()
	s.debug()
	s.invoker()
	s.initHttpServer()
	s.initRouter()
	s.initValidator()
	s.govern()
	return s.err
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	go func() {
		<-stopCh
		s.Add(1)
		xdefer.Clean()
		s.Done()
	}()
	xdefer.Register(func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		xconsole.Red("http server shutdown")
		return s.Server.Shutdown(ctx)
	})
	xconsole.Greenf("Start listening on:", s.Server.Addr)
	if err = s.Server.ListenAndServe(); err == http.ErrServerClosed {
		return nil
	}
	s.Wait()
	return err
}

func (s *Server) debug() {
	xconsole.ResetDebug(xapp.Debug())
	xapp.PrintVersion()
}

func (s *Server) initCfg() {
	if s.err != nil {
		return
	}
	var data xcfg.DataSource
	data, s.err = manager.NewDataSource(xflag.NString("run", "xcfg"))
	if s.err != nil {
		return
	}
	s.err = xcfg.LoadFromDataSource(data, toml.Unmarshal)
}

func (s *Server) invoker() {
	if s.err != nil {
		return
	}
	xdefer.Register(func() error {
		return xinvoker.Close()
	})
	xinvoker.Register(
		//xgorm.Register("mysql"),
		//xredis.Register("redis"),
	)
	s.err = xinvoker.Init()
}

func (s *Server) initHttpServer() {
	if s.err != nil {
		return
	}
	s.Server = new(http.Server)
	s.Server.Addr = xcfg.GetString("server.addr")
}

func (s *Server) initRouter() {
	if s.err != nil {
		return
	}
	s.Server.Handler = registry.Engine()
}

func (s *Server) initValidator() {
	if s.err != nil {
		return
	}
	s.err = xvalidator.Init(xcfg.GetString("server.locale"), myValidator.RegisterValidation)
}

func (s *Server) govern() {
	if s.err != nil {
		return
	}
	xmonitor.Run()
	xtrace.Init("trace.jaeger")
	go xgovern.Run()
}
