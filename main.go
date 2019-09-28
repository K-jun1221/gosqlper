package main

import (
	"fmt"
	"gosqlper/db"
	"gosqlper/lib"
)

type IEvent struct {
	EventID   string `json:"event_id" db:"event_id"`
	EventName string `json:"event_name" db:"event_name"`

	// StartDate           string `json:"start_date" db:"start_date"`
	// EndDate             string `json:"end_date" db:"end_date"`
	// Location            string `json:"location" db:"location"`
	// TargetUser          string `json:"target_user" db:"target_user"`
	// CreatedUserID       string `json:"created_user_id`
	// ParticipantLimitNum int64  `json:"participant_limit_num"`
	// EventDetail         string `json:"event_detail"`
}

func main() {
	db.Initialize("hidakkathon:hidakkathon@tcp(127.0.0.1:3306)/sugori_rendez_vous")

	var obj IEvent
	sql := lib.SelectSQL{
		Select: []string{"event_id", "event_name"},
		From:   "i_event",
		Where:  "event_id = 1",
	}

	lib.SelectRow(db.DB, sql, &obj)
	fmt.Println("obj event_id: ", obj.EventID)
	fmt.Println("obj event_name: ", obj.EventName)
}
