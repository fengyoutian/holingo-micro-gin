package dao

import (
	"context"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Provider = wire.NewSet(New, NewClient)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error)
	AddArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error)
	SearchArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error)
}

// dao dao.
type dao struct {
	client *client
}

// New new a dao and return.
func New(c *client) (d Dao, cf func(), err error) {
	return newDao(c)
}

func newDao(c *client) (d *dao, cf func(), err error) {
	d = &dao{
		client: c,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {

}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

// SayHello
func (d *dao) SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error) {
	res, err := d.client.HolingoService.SayHello(ctx, req)
	logrus.Infof("dao.SayHello res(%s)", res.Content)
	if err != nil {
		err = errors.Wrapf(err, "%v", req.Name)
		return
	}
	*reply = *res
	logrus.Infof("dao.SayHello reply(%s)", reply.Content)
	return
}

// AddArticle
func (d *dao) AddArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	res, err := d.client.HolingoService.AddArticle(ctx, req)
	if err != nil {
		err = errors.Wrapf(err, "%v", req)
	} else {
		*reply = *res
	}
	logrus.Infof("dao.AddArticle reply(%v)", *reply)
	return
}

// SearchArticle
func (d *dao) SearchArticle(ctx context.Context, req *holingo.Article, reply *holingo.Article) (err error) {
	res, err := d.client.HolingoService.SearchArticle(ctx, req)
	if err != nil {
		err = errors.Wrapf(err, "%v", req)
	} else {
		*reply = *res
	}
	logrus.Infof("dao.SearchArticle reply(%v)", *reply)
	return
}
