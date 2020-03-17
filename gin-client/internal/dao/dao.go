package dao

import (
	"context"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var Provider = wire.NewSet(New, NewClient)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	SayHello(ctx context.Context, req *holingo.HelloReq, reply *holingo.HelloResp) (err error)
}

// dao dao.
type dao struct {
	client holingo.HolingoService
}

// New new a dao and return.
func New(c holingo.HolingoService) (d Dao, cf func(), err error) {
	return newDao(c)
}

func newDao(c holingo.HolingoService) (d *dao, cf func(), err error) {
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
	if reply, err = d.client.SayHello(ctx, req); err != nil {
		err = errors.Wrapf(err, "%v", req.Name)
	}
	return
}
