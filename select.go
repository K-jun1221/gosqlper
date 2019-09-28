package gosqlper

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SelectRow(sql SelectSQL) error {
	rawSql, err := sql.MakeSQL()
	if err != nil {
		return errors.New("failed to parse to raw sql")
	}

	_, err = GetDB()
	if err != nil {
		return err
	}

	fmt.Println("Will call QueryRow", rawSql)
	// TODO
	// obj := ""
	// err = mysql.QueryRow(rawSql).Scan(obj)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func SelectRows(sql SelectSQL) error {
	rawSql, err := sql.MakeSQL()
	if err != nil {
		return errors.New("failed to parse to raw sql")
	}

	_, err = GetDB()
	if err != nil {
		return err
	}

	fmt.Println("Will call Query", rawSql)
	return nil
}
