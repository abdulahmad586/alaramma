package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

// DB represents the database connection.
var DB *sql.DB

// InitDB initializes the database connection.
func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
	

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
}

// CloseDB closes the database connection.
func CloseDB() {
    if err := DB.Close(); err != nil {
        log.Fatal(err)
    }
}
