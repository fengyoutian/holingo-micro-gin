// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/dao"
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/server/http"
	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/service"
	"github.com/google/wire"
)

//go:generate wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, NewApp))
}
