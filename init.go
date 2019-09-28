package gosqlper

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize @params: dsn, @response error
func Initialize(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

// GetDB @params: None, @response *sql.DB, error
func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("DB was not initialized")
	}

	return db, nil
}
