package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB  // Change db to DB to export it

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to the global DB

	if err != nil {
		panic("Failed to connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTable()
}

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		date DATETIME NOT NULL,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id INTEGER NOT NULL
	);`

	_, err := DB.Exec(createEventsTable) // Use DB
	if err != nil {
		panic("Failed to create events table")
	}
}
