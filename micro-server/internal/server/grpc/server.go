package grpc

import (
	"time"

	"github.com/fengyoutian/holingo-micro-gin/pkg/config"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/service"
	"github.com/fengyoutian/holingo-micro-gin/tool"
	"github.com/fengyoutian/holingo-util/file"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/sirupsen/logrus"
)

func New(service *service.Service) (s *micro.Service, err error) {
	var (
		serverCfg config.GrpcConfig
		etdcCfg   config.EtcdConfig
		y         *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("grpc.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("server", &serverCfg); err != nil {
		return
	}
	if err = y.Unmarshal("etcd", &etdcCfg); err != nil {
		return
	}
	logrus.Infof("go-micro: %v, etcd: %v \n ", serverCfg, etdcCfg)

	// Register EtcdConfig
	registry := etcd.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = etdcCfg.Addrs
		opt.Timeout = etdcCfg.Timeout * time.Second
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
				logrus.Infof("Environment set to %s", env)
			}
			return nil
		}),
	)

	// Register Handler
	holingo.RegisterHolingoHandler(engine.Server(), service)

	// Run engine
	if err = engine.Run(); err != nil {
		logrus.Fatal(err)
	}
	return
}
