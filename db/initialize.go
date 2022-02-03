package db

import (
	"database/sql"
	"time"

	"upv.life/server/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Initialize() {
	drv, err := sql.Open("mysql", config.MysqlDsn)
	db = sqlx.NewDb(drv, "mysql")

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
}
