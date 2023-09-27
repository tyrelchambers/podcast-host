package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var DbUrl = getEnv("DB_URL")

const (
	host           = "localhost"
	port           = "5432"
	user           = "dbadmin"
	password       = "JUYVgv3vutcr4hjJd"
	dbname         = "db"
	BUNNY_URL_BASE = "https://podcast-files.b-cdn.net"
)
