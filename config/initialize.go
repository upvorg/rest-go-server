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

	SMMSUserName string
	SMMSPassword string

	QQAppID  string
	QQAppKey string
)

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	if os.Getenv("APP_MODE") != "" {
		AppMode = os.Getenv("APP_MODE")
	}

	if os.Getenv("APP_PORT") != "" {
		AppPort = os.Getenv("APP_PORT")
	}

	if os.Getenv("APP_DOMAIN") != "" {
		Domain = os.Getenv("APP_DOMAIN")
	}

	MysqlDsn = os.Getenv("MYSQL_DSN")
	JwtSalt = os.Getenv("JWT_SALT")

	SMMSUserName = os.Getenv("SMMS_USERNAME")
	SMMSPassword = os.Getenv("SMMS_PASSWORD")

	QQAppID = os.Getenv("QQ_APP_ID")
	QQAppKey = os.Getenv("QQ_APP_KEY")
}
