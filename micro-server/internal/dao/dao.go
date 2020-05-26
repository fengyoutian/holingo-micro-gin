package dao

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/logger"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/model"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Provider = wire.NewSet(New, NewDB)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	AddArticle(ctx context.Context, req *model.Article) (reply *model.Article, err error)
	SearchArticle(ctx context.Context, req *model.Article) (reply *model.Article, err error)
}

// dao dao.
type dao struct {
	db *gorm.DB
}

// New new a dao and return.
func New(db *gorm.DB) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *gorm.DB) (d *dao, cf func(), err error) {
	d = &dao{
		db: db,
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

// AddArticle add article to db
func (d *dao) AddArticle(ctx context.Context, req *model.Article) (reply *model.Article, err error) {
	//exec := d.db.Exec(
	//	fmt.Sprintf("INSERT INTO %s (author, title, content) VALUES (?,?,?)", req.TableName()),
	//	req.Author, req.Title, req.Content)
	exec := d.db.Table(req.TableName()).Create(&req)
	if err = exec.Error; err != nil {
		return
	}
	return req, nil
}

// SearchArticle search article for db
func (d *dao) SearchArticle(ctx context.Context, req *model.Article) (reply *model.Article, err error) {
	err = d.db.Raw(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", req.TableName()), req.ID).Scan(req).Error
	if err != nil {
		logger.Errorf("dao.SearchArticle Error(%v)", err)
	}
	return req, err
}
