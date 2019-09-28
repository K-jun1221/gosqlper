package gosqlper

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SelectRow(sql SelectSQL) {
	// db.Query("SELECT")
	fmt.Println("SelectRow")
}

func SelectRows(sql SelectSQL) {
	// db.Query("SELECT")
	fmt.Println("SelectRows")
}
