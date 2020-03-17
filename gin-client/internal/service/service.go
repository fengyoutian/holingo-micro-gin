package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/dao"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(holingo.HolingoHandler), new(*Service)))

// Service service.
type Service struct {
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty, reply *empty.Empty) (err error) {
	return s.dao.Ping(ctx)
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error) {
	reply = new(holingo.HelloResp)
	err = s.dao.SayHello(ctx, req, reply)
	logrus.Infof("client.SayHello() reply %s \n", reply.Content)
	return
}
