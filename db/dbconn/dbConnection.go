package dbconn

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

// New Instantiates Postgres service.
func NewPostgres() (*Postgres, error) {
	db, err := getDBInstance(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func getDBInstance(dbKey string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbKey)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
