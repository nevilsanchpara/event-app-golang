package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "123"
    dbname   = "event"
)

func InitDB() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    createTables()
}

func createTables() {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    )
    `

    _, err := DB.Exec(createUsersTable)
    if err != nil {
        panic(fmt.Errorf("failed to create users table: %v", err))
    }

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP NOT NULL,
        user_id INTEGER REFERENCES users(id)
    )
    `

    _, err = DB.Exec(createEventsTable)
    if err != nil {
        panic(fmt.Errorf("failed to create events table: %v", err))
    }

    createRegistrationsTable := `
    CREATE TABLE IF NOT EXISTS registrations (
        id SERIAL PRIMARY KEY,
        event_id INTEGER REFERENCES events(id),
        user_id INTEGER REFERENCES users(id)
    )
    `

    _, err = DB.Exec(createRegistrationsTable)
    if err != nil {
        panic(fmt.Errorf("failed to create registrations table: %v", err))
    }
}
