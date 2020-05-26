package dao

import (
	"time"

	"github.com/micro/go-micro/v2/config"

	"github.com/micro/go-micro/v2/logger"

	"github.com/fengyoutian/holingo-micro-gin/micro-server/internal/model"

	myConfig "github.com/fengyoutian/holingo-micro-gin/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg myConfig.DBConfig
	)
	// config load on main.go
	if err = config.Get("hosts", "database").Scan(&cfg); err != nil {
		return
	}
	logger.Infof("db: %v\n ", cfg)

	db, err = gorm.Open(cfg.Dialect, cfg.DSN)
	if err != nil {
		logger.Errorf("db.Open Error(%v)", err)
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
