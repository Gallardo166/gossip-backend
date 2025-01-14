package initializers

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}

	connStr := os.Getenv("CONNSTR")
	var dbErr error
	DB, dbErr = sqlx.Connect("postgres", connStr)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %s", dbErr)
	}
}
