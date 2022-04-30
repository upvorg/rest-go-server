package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppMode  string = "debug"
	AppPort  string = "8080"
	Domain   string = "localhost"
	MysqlDsn string
	JwtSalt  string
)

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	if os.Getenv("ENV") != "" {
		AppMode = os.Getenv("ENV")
	}
	if os.Getenv("PORT") != "" {
		AppPort = os.Getenv("APP_PORT")
	}
	if os.Getenv("APP_DOMAIN") != "" {
		Domain = os.Getenv("APP_DOMAIN")
	}

	MysqlDsn = os.Getenv("MYSQL_DSN")
	JwtSalt = os.Getenv("JWT_SALT")
}
