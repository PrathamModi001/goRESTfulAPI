package models

import "time"

// all the logic to storing event data in the database

// Event struct: Defines the shape of the event
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       // to link the event to the user
}

var events = []Event{}

func (e *Event) Save() error {
	// add it to the database
	events = append(events, *e)
	return nil
}

func GetAllEvents() []Event {
	return events
}
