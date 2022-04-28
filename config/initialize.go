package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppMode  string = "debug"
	AppPort  string
	MysqlDsn string
	JwtSalt  string
	Domain   string
)

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	if os.Getenv("ENV") != "" {
		AppMode = os.Getenv("ENV")
	}

	AppPort = os.Getenv("APP_PORT")
	MysqlDsn = os.Getenv("MYSQL_DSN")
	JwtSalt = os.Getenv("JWT_SALT")
	Domain = os.Getenv("APP_DOMAIN")
}
