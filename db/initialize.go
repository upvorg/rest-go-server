package db

import (
	"database/sql"
	"time"

	"upv.life/server/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Sqlx *sqlx.DB
)

func Initialize() {
	drv, err := sql.Open("mysql", config.MysqlDsn)
	Sqlx = sqlx.NewDb(drv, "mysql")

	Sqlx.SetMaxIdleConns(10)
	Sqlx.SetMaxOpenConns(100)
	Sqlx.SetConnMaxLifetime(time.Hour)
	defer Sqlx.Close()

	if err != nil {
		panic(err.Error())
	}
}
