package di

import (
	"context"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/logger"

	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/service"
)

type App struct {
	svc  *service.Service
	http *http.Server
}

func NewApp(svc *service.Service, h *http.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := h.Shutdown(ctx); err != nil {
			logger.Errorf("httpSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
