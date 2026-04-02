package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	DataTIme    time.Time
	Location    string `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, datetime, location, user_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Could not prepare statement.")

	}
	defer stmt.Close()
	results, err := stmt.Exec(e.Name, e.Description, e.DataTIme, e.Location, e.UserID)
	if err != nil {
		panic("Could not execute statement.")
	}
	id, err := results.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		panic("Could not query database.")
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTIme, &event.UserID)
		if err != nil {
			panic("Could not scan row.")
		}
		events = append(events, event)
	}
	return events, nil
}

func GetOneEvent(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	var event Event
	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTIme, &event.UserID)
	if err != nil {
		panic("Could not scan row.")
	}

	return event, nil

}
