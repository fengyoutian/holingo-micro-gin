package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fengyoutian/holingo-micro-gin/tool"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/di"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	tool.Config.Init()
	logrus.Info("micro-server start")

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logrus.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			logrus.Info("micro-server exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
