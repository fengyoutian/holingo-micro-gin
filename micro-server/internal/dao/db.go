package dao

import (
	"time"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/model"

	"github.com/fengyoutian/holingo-micro-gin/pkg/config"
	"github.com/fengyoutian/holingo-micro-gin/tool"
	"github.com/fengyoutian/holingo-util/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg config.DBConfig
		y   *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("db.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("client", &cfg); err != nil {
		return
	}
	logrus.Infof("db: %v\n ", cfg)

	db, err = gorm.Open(cfg.Dialect, cfg.DSN)
	if err != nil {
		logrus.Errorf("db.Open Error(%v)", err)
	}
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	db.DB().SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Second)

	err = db.AutoMigrate(&model.Article{}).Error
	cf = func() {
		db.Close()
	}
	return
}
