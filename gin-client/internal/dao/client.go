package dao

import (
	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/fengyoutian/holingo-micro-gin/tool"
	"github.com/fengyoutian/holingo-util/file"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
)

type ClientConfig struct {
	Name    string
	Version string
	Dial    string
}

// NewClient new grpc client
func NewClient() (svc holingo.HolingoService, err error) {
	var (
		cfg ClientConfig
		y   *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("grpc.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("client", &cfg); err != nil {
		return
	}
	logrus.Infof("cfg: %s\n ", cfg)

	// New Service
	engine := micro.NewService(
		micro.Name(cfg.Name),
		micro.Version(cfg.Version),
		micro.Address(cfg.Dial),
	)

	svc = holingo.NewHolingoService(cfg.Name, engine.Client())
	return
}
