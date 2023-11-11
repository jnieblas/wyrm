package driver

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "testdb.db")

	if err != nil {
		log.Println(err)
		log.Fatal("Could not open database")
	}

	return db
}
