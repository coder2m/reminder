package xrpc

import (
	"fmt"
	xapp "github.com/coder2m/component"
	"github.com/coder2m/component/pkg/xdefer"
	"github.com/coder2m/component/pkg/xnet"
	"github.com/coder2m/component/xgrpc"
	serverinterceptors "github.com/coder2m/component/xgrpc/server"
	"github.com/coder2m/component/xregistry"
	"github.com/coder2m/component/xregistry/xetcd"
	"github.com/coder2m/reminder/pkg/constant"
	"google.golang.org/grpc"
	"time"
)

type Config struct {
	EtcdAddr         string        `mapStructure:"etcd_addr"`
	ServerIp         string        `mapStructure:"ip"`
	ServerPort       int           `mapStructure:"port"`
	RegisterTTL      time.Duration `mapStructure:"register_ttl"`
	RegisterInterval time.Duration `mapStructure:"register_interval"`
	Timeout          time.Duration `mapStructure:"timeout"`
}

func DefaultConfig() *Config {
	host, port, err := xnet.GetLocalMainIP()
	if err != nil {
		host = "localhost"
	}
	return &Config{
		EtcdAddr:         "127.0.0.1:2379",
		ServerIp:         host,
		ServerPort:       port,
		RegisterTTL:      30 * time.Second,
		RegisterInterval: 15 * time.Second,
		Timeout:          30 * time.Second,
	}
}

func (c Config) Addr() string {
	return fmt.Sprintf("%v:%v", c.ServerIp, c.ServerPort)
}

func DefaultOption(c *Config) []grpc.ServerOption {
	return []grpc.ServerOption{
		xgrpc.WithUnaryServerInterceptors(
			serverinterceptors.CrashUnaryServerInterceptor(),
			serverinterceptors.PrometheusUnaryServerInterceptor(),
			serverinterceptors.XTimeoutUnaryServerInterceptor(c.Timeout),
			serverinterceptors.TraceUnaryServerInterceptor(),
		),
		xgrpc.WithStreamServerInterceptors(
			serverinterceptors.CrashStreamServerInterceptor(),
			serverinterceptors.PrometheusStreamServerInterceptor(),
		),
	}
}

func DefaultRegistryEtcd(c *Config) (err error) {
	var etcdR xregistry.Registry
	conf := xetcd.EtcdV3Cfg{
		Endpoints: []string{c.EtcdAddr},
	}
	etcdR, err = xetcd.NewRegistry(conf) //注册
	if err != nil {
		return
	}

	etcdR.Register(
		xregistry.ServiceName(xapp.Name()),
		xregistry.ServiceNamespaces(constant.DefaultNamespaces),
		xregistry.Address(c.Addr()),
		xregistry.RegisterTTL(c.RegisterTTL),
		xregistry.RegisterInterval(c.RegisterInterval),
	)

	xdefer.Register(func() error {
		etcdR.Close()
		return nil
	})
	return
}
