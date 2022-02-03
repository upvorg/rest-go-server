package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MysqlDsn         string
	JwtSlat          string
	UserPasswordSalt string
	AppPort          string
)

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	MysqlDsn = os.Getenv("MYSQL_DSN")
	JwtSlat = os.Getenv("JWT_SLAT")
	UserPasswordSalt = os.Getenv("USER_PASSWORD_SALT")
	AppPort = os.Getenv("APP_PORT")
}
