package models

import (
	"time"
	"example.com/restAPI/db"
)

// all the logic to storing event data in the database

// Event struct: Defines the shape of the event
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       // to link the event to the user
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, date, location, description, user_id)
	VALUES (?, ?, ?, ?, ?)
	` // ? is a placeholder for the actual values to prevent SQL injection

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer will run the stmt.Close() after the function returns
	defer stmt.Close()

	// safe way to inject the values
	result, err := stmt.Exec(e.Name, e.DateTime, e.Location, e.Description, e.UserID) 
	if err != nil {
		return err
	}

	// we still have not saved the correct ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() []Event {
	return events
}
