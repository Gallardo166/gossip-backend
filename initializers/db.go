package initializers

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	connStr := os.Getenv("CONNSTR")
	var dbErr error
	DB, dbErr = sqlx.Connect("postgres", connStr)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %s", dbErr)
	}
}
