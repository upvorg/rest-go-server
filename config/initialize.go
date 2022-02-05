package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppMode          string
	AppPort          string
	MysqlDsn         string
	JwtSalt          string
	UserPasswordSalt string
)

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	AppMode = os.Getenv("APP_MODE")
	AppPort = os.Getenv("APP_PORT")
	MysqlDsn = os.Getenv("MYSQL_DSN")
	JwtSalt = os.Getenv("JWT_SALT")
	UserPasswordSalt = os.Getenv("USER_PASSWORD_SALT")
}
