package grpc

import (
	"time"

	myConfig "github.com/fengyoutian/holingo-micro-gin/pkg/config"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/service"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func New(service *service.Service) (s *micro.Service, err error) {
	var (
		serverCfg myConfig.GrpcConfig
		etcdCfg   myConfig.EtcdConfig
	)
	// config load on main.go
	if err = config.Get("hosts", "server").Scan(&serverCfg); err != nil {
		return
	}
	if err = config.Get("hosts", "etcd").Scan(&etcdCfg); err != nil {
		return
	}
	logger.Infof("go-micro: %v, etcd: %v \n ", serverCfg, etcdCfg)

	// Register EtcdConfig
	registry := etcd.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = etcdCfg.Addrs
		opt.Timeout = etcdCfg.Timeout * time.Second
	})

	// New Service
	engine := micro.NewService(
		micro.Name(serverCfg.Name),
		micro.Version(serverCfg.Version),
		micro.Address(serverCfg.Addr),

		micro.Registry(registry),
	)
	s = &engine

	// Initialise engine
	engine.Init(
		micro.Action(func(c *cli.Context) error {
			env := c.String(serverCfg.Name)
			if len(env) > 0 {
				logger.Infof("Environment set to %s", env)
			}
			return nil
		}),
	)

	// Register Handler
	holingo.RegisterHolingoHandler(engine.Server(), service)

	// Run engine
	if err = engine.Run(); err != nil {
		logger.Fatal(err)
	}
	return
}
