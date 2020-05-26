package service

import (
	"context"

	"github.com/fengyoutian/holingo-micro-gin/gin-client/internal/dao"
	"github.com/micro/go-micro/v2/logger"

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
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty, reply *empty.Empty) (err error) {
	return s.dao.Ping(ctx)
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error) {
	err = s.dao.SayHello(ctx, req, reply)
	logger.Infof("service.SayHello reply %s \n", (*reply).Content)
	return
}

// AddArticle grpc demo func
func (s *Service) AddArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	err = s.dao.AddArticle(ctx, req, reply)
	logger.Infof("service.AddArticle reply(%v) \n", reply)
	return
}

// SearchArticle grpc demo func
func (s *Service) SearchArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	err = s.dao.SearchArticle(ctx, req, reply)
	logger.Infof("service.SearchArticle reply(%v) \n", reply)
	return
}
