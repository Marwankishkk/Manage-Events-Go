package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
	return DB, nil
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL);`
	_, err := DB.Exec(createUsersTable)

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER ,
			foreign key (user_id) references users(id)
		);`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create users table.")
	}
	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			event_id INTEGER,
			foreign key (user_id) references users(id),
			foreign key (event_id) references events(id),
			UNIQUE (user_id, event_id)
		);`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table.")
	}

}
