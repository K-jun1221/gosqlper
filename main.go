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

	// var obj IEvent
	sql := lib.SelectSQL{
		Select: []string{"event_id", "event_name"},
		From:   "i_event",
	}

	var obj IEvent
	// var objs []interface{}
	lib.QueryRow(db.DB, sql, &obj)
	fmt.Println("obj event_id: ", obj.EventID)
	fmt.Println("obj event_name: ", obj.EventName)

	// err := lib.Query(db.DB, sql, &obj, objs)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("result: ", objs)
}
