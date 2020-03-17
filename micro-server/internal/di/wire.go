// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/server/grpc"
	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/service"
	"github.com/google/wire"
)

//go:generate wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(service.Provider, grpc.New, NewApp))
}
