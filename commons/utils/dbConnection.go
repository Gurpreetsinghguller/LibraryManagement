package utils

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	DB *sqlx.DB
}

// New Instantiates Postgres service.
func NewPostgres() (*Postgres, error) {
	db, err := getDBInstance(os.Getenv(""))
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func getDBInstance(envDbKey string) (*sqlx.DB, error) {
	dsn, retrieved := os.LookupEnv(envDbKey)
	if !retrieved {
		return nil, fmt.Errorf("error retrieving database connection string from params")
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
