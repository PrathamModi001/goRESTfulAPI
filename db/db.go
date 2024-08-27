package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Connect to the SQLite database

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10) // Set maximum number of open connections
	DB.SetMaxIdleConns(5)  // Set maximum number of idle connections

	createTables() // Create tables if they do not exist
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL,                   
		description TEXT NOT NULL,            
		location TEXT NOT NULL,               
		dateTime DATETIME NOT NULL,           
		user_id INTEGER                       
	)
	`

	_, err := DB.Exec(createEventsTable) // Execute the table creation SQL command

	if err != nil {
		panic("Could not create events table.")
	}

}
