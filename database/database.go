/*
Package database
holds query functions
*/
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./cars.db")
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("failed to ping db", err)
	}
}
