package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"upv.life/server/config"
	"upv.life/server/model"
)

var Orm *gorm.DB

func Initialize() {
	logLevel := logger.Silent
	if config.AppMode == "debug" {
		logLevel = logger.Info
	}
	gorm, err := gorm.Open(mysql.Open(config.MysqlDsn),
		&gorm.Config{Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: false,
				Colorful:                  false,
			},
		)})
	if err != nil {
		panic(err.Error())
	}

	Orm = gorm
	sqlDB, _ := gorm.DB()
	gorm.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.VideoMeta{},
		&model.PostRanking{},
		&model.Video{},
		&model.Comment{},
		&model.Like{},
		&model.Collection{},
		&model.Feedback{},
	)

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
