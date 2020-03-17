package di

import (
	"context"
	"time"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/service"

	"github.com/micro/go-micro/v2"
)

type App struct {
	svc  *service.Service
	grpc *micro.Service
}

func NewApp(svc *service.Service, g *micro.Service) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		grpc: g,
	}
	closeFunc = func() {
		_, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		cancel()
	}
	return
}
