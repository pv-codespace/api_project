package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		log.Fatalf("Could not create table users: %v", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create table event: %v", err)
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registration(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		log.Fatalf("Could not create table registration: %v", err)
	}

	fmt.Println("DB connected and table created")
}
