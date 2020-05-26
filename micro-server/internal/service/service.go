package service

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/model"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/dao"

	"github.com/google/wire"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/golang/protobuf/ptypes/empty"
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
func (s *Service) Ping(ctx context.Context, e *empty.Empty, r *empty.Empty) error {
	r = new(empty.Empty)
	logger.Info("grpc.Ping() pong. \n")
	return nil
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error) {
	reply.Content = "hello word"
	logger.Infof("grpc.SayHello(%s) reply(%s) \n", req.Name, reply.Content)
	return
}

// AddArticle gorm demo func
func (s *Service) AddArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	article, err := s.dao.AddArticle(ctx, &model.Article{
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return
	}
	articleHandle(reply, article)
	logger.Infof("service.AddArticle(%s) reply(%v)", req.Title, reply)
	return
}

// SearchArticle gorm demo func
func (s *Service) SearchArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	article, err := s.dao.SearchArticle(ctx, &model.Article{ID: req.Id})
	if err != nil {
		return
	}
	articleHandle(reply, article)
	logger.Infof("service.SearchArticle(%d) reply(%v)", req.Id, reply)
	return
}

func articleHandle(reply *holingo.Article, article *model.Article) {
	reply.Id = article.ID
	reply.Author = article.Author
	reply.Title = article.Title
	reply.Content = article.Content
	reply.ModifyTime = article.ModifyTime
	reply.CreateTime = article.CreateTime
}
