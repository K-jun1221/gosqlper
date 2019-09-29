package main

import (
	"fmt"
	"gosqlper/db"
	"gosqlper/lib"
)

type IEvent struct {
	EventID   string `json:"event_id" db:"event_id"`
	EventName string `json:"event_name" db:"event_name"`
}

func main() {
	db.Initialize("hidakkathon:hidakkathon@tcp(127.0.0.1:3306)/sugori_rendez_vous")

	var obj IEvent
	var objs []IEvent
	sql := lib.SelectSQL{
		Select: []string{"event_id", "event_name"},
		From:   "i_event",
	}

	err := lib.QueryRow(db.DB, sql, &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("obj: ", obj)

	err = lib.Query(db.DB, sql, &objs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("objs: ", objs)

	sql2 := &lib.InsertSQL{
		Into:   "i_event_tag (event_id, tag_id)",
		Values: "(3, 3)",
	}

	_, err = lib.Exec(db.DB, sql2)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}
