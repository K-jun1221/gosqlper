package main

import (
	"database/sql"
	"fmt"

	"github.com/K-jun1221/gosqlper"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"create_date"`
	UpdatedAt string `db:"update_date"`
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sample"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	sql := "SELECT name, id, create_date, update_date FROM todo where id = 3"

	var todo Todo
	err = gosqlper.QueryRow(db, sql, &todo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(todo)


	sql2 := "SELECT name, id, create_date, update_date FROM todo"

	var todos []Todo
	err = gosqlper.Query(db, sql2, &todos)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(todos)
}
