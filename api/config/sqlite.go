package config

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func GetDbConnection() (*sql.DB, error) {
	dbPath := "./db/main.db"
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Printf("sqlite opening error: %s", err)
		return nil, err
	}
	return db, nil
}
