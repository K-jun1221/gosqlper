package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	fmt.Println("connecting to mysql db: ", dsn)
	if err != nil {
		return err
	}

	return nil
}
