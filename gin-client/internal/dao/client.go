package dao

import (
	"time"

	"github.com/micro/go-micro/v2/config"

	"github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2/registry"

	myConfig "github.com/fengyoutian/holingo-micro-gin/pkg/config"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// dao dao.
type client struct {
	HolingoService holingo.HolingoService
}

var engine micro.Service

// NewClient new grpc client
func NewClient() (c *client, cf func(), err error) {
	var (
		clientCfg myConfig.GrpcConfig
		etdcCfg   myConfig.EtcdConfig
	)
	// config load on main.go
	if err = config.Get("hosts", "client").Scan(&clientCfg); err != nil {
		return
	}
	if err = config.Get("hosts", "etcd").Scan(&etdcCfg); err != nil {
		return
	}
	logger.Infof("client: %v, etcd: %v \n ", clientCfg, etdcCfg)

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
		logger.Errorf("client.Close.Server Error(%v)", err)
	}
}
