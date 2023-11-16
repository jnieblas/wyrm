package migration

import (
	"database/sql"
	"log"
)

// This only needs to be called on install / version upgrade
func ProvisionDB() {
	log.SetPrefix("MIGRATION: ")
	migration1_createScriptsTable()
}

func migration1_createScriptsTable() {
	db, err := sql.Open("sqlite3", "testdb.db")

	if err != nil {
		log.Println(err)
		log.Fatal("Could not open database")
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scripts (
			name TEXT NOT NULL PRIMARY KEY,
			path TEXT NOT NULL,
			command TEXT NOT NULL,
			description TEXT
		)
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Unable to create the scripts table.")
	}
}
