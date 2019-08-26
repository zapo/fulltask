package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDatabase(dsn string) {
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Couldnt parse DSN")
	}
}
