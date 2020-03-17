// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/dao"
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/server/http"
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	holingoService, err := dao.NewClient()
	if err != nil {
		return nil, nil, err
	}
	daoDao, cleanup, err := dao.New(holingoService)
	if err != nil {
		return nil, nil, err
	}
	serviceService, cleanup2, err := service.New(daoDao)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	server, err := http.New(serviceService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup3, err := NewApp(serviceService, server)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}