package models

import (
	"time"

	"example.com/rest-api/db"
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

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	// ? is a placeholder for the actual values to prevent SQL injection attacks

	// Prepare the query 
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer will run the stmt.Close() after the function returns
	defer stmt.Close()

	// safe way to inject the values
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// we still have not saved the correct ID
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// db package DB is a pointer to the database connection in db.go (global variable)
	// you could have prepared and exec the query here, but this is a much simpler query
	// could also have used Exec directly
	// we use Query when we expect multiple rows (USED FOR FETCHING DATA)
	// we use Exec when we expect a CHANGE in the database (USED FOR CHANGING DB)
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create a slice of events
	var events []Event

	// iterate over the rows
	for rows.Next() {
		var event Event

		// give the address of pointers IN THE ORDER of the columns
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
