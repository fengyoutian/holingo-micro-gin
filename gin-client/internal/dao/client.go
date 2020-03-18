package dao

import (
	"time"

	"github.com/micro/go-micro/v2/registry"

	"github.com/fengyoutian/holingo-micro-gin/pkg/config"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/fengyoutian/holingo-micro-gin/tool"
	"github.com/fengyoutian/holingo-util/file"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/sirupsen/logrus"
)

// dao dao.
type client struct {
	HolingoService holingo.HolingoService
}

var engine micro.Service

// NewClient new grpc client
func NewClient() (c *client, cf func(), err error) {
	var (
		clientCfg config.GrpcConfig
		etdcCfg   config.EtcdConfig
		y         *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("grpc.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("client", &clientCfg); err != nil {
		return
	}
	if err = y.Unmarshal("etcd", &etdcCfg); err != nil {
		return
	}
	logrus.Infof("client: %v, etcd: %v \n ", clientCfg, etdcCfg)

	// Register EtcdConfig
	registry := etcd.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = etdcCfg.Addrs
		opt.Timeout = etdcCfg.Timeout * time.Second
	})

	// New Service
	engine = micro.NewService(
		micro.Name(clientCfg.Name),
		micro.Version(clientCfg.Version),
		//micro.Address(clientCfg.Addr),

		micro.Registry(registry),
	)

	c = &client{
		HolingoService: holingo.NewHolingoService(clientCfg.Name, engine.Client()),
	}
	cf = c.Close
	return
}

// Close close the resource.
func (c *client) Close() {
	if err := engine.Server().Stop(); err != nil {
		logrus.Errorf("client.Close.Server Error(%v)", err)
	}
}
