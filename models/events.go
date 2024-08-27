package models

import (
	"example.com/restAPI/db"
	"time"
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

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	// db package DB is a pointer to the database connection in db.go (global variable)
	// you could have prepared and exec the query here, but this is a much simpler query
	// could also have used Exec directly
	// we use Query when we expect multiple rows (USED FOR FETCHING DATA)
	// we use Exec when we expect a CHANGE in the database (USED FOR CHANGING DB)
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	// create a slice of events
	events := []Event{}

	// iterate over the rows
	for rows.Next() {
		var event Event
		var dateTimeStr string

		// give the address of pointers IN THE ORDER of the columns
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
		if err != nil {
			panic(err)
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (Event, error) {
    query := `SELECT * FROM events WHERE id = ?`
    row := db.DB.QueryRow(query, id)

    var event Event
    var dateTimeStr string
    err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
    if err != nil {
        return Event{}, err
    }

    event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
    if err != nil {
        return Event{}, err
    }

    return event, nil
}

