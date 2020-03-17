package di

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

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
			logrus.Errorf("httpSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
