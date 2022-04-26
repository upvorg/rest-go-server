package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"upv.life/server/config"
)

var Orm *gorm.DB

func Initialize() {
	gorm, err := gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	Orm = gorm
	sqlDB, err := gorm.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
