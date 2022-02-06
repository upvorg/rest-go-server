package db

import (
	"database/sql"
	"time"

	"upv.life/server/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Initialize() {
	drv, err := sql.Open("mysql", config.MysqlDsn)
	if err != nil {
		panic(err.Error())
	}

	DB = sqlx.NewDb(drv, "mysql")
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)
}
