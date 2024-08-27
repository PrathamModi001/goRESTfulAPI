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
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	_ ,err := DB.Exec(createUsersTable) // Execute the table creation SQL command
	if err != nil {
		panic("Could not create users table.")
	}


	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL,                   
		description TEXT NOT NULL,            
		location TEXT NOT NULL,               
		dateTime DATETIME NOT NULL,           
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`// establish connection between the user and the event by FK

	_, err = DB.Exec(createEventsTable) // Execute the table creation SQL command
	if err != nil {
		panic("Could not create events table.")
	}
}
