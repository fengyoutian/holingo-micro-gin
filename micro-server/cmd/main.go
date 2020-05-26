package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"

	"github.com/micro/go-micro/v2/logger"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/di"
)

func main() {
	flag.Parse()
	// load config file
	if err := config.Load(file.NewSource(file.WithPath("../configs/config.yaml"))); err != nil {
		return
	}
	logger.Info("micro-server start")

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logger.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			logger.Info("micro-server exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
