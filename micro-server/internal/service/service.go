package service

import (
	"context"

	"github.com/google/wire"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/sirupsen/logrus"

	"github.com/golang/protobuf/ptypes/empty"
)

var Provider = wire.NewSet(New, wire.Bind(new(holingo.HolingoHandler), new(*Service)))

// Service service.
type Service struct {
}

// New new a service and return.
func New() (s *Service, cf func(), err error) {
	s = &Service{}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty, r *empty.Empty) error {
	r = new(empty.Empty)
	logrus.Info("grpc.Ping() pong. \n")
	return nil
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error) {
	reply = &holingo.HelloResp{
		Content:              "hello word!",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	logrus.Infof("grpc.SayHello(%s) reply(%s) \n", req.Name, reply.Content)
	return
}
