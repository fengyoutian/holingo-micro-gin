package grpc

import (
	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/service"
	"github.com/fengyoutian/holingo-micro-gin/tool"
	"github.com/fengyoutian/holingo-util/file"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Name    string
	Version string
	Addr    string
}

func New(service *service.Service) (s *micro.Service, err error) {
	var (
		cfg ServerConfig
		y   *file.YAML
	)
	if y, err = file.Load(tool.Cofig.GetConfigPath("grpc.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("server", &cfg); err != nil {
		return
	}
	logrus.Infof("cfg: %s\n ", cfg)

	// New Service
	engine := micro.NewService(
		micro.Name(cfg.Name),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr),
	)
	s = &engine

	// Initialise engine
	engine.Init(
		micro.Action(func(c *cli.Context) error {
			env := c.String(cfg.Name)
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

func registerHandler(srv *micro.Service) {

}
