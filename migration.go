package main

import (
	"log"
)

// This only needs to be called on install / version upgrade
func provisionDB() {
	log.SetPrefix("MIGRATION: ")
	migration1_createScriptsTable()
}

func migration1_createScriptsTable() {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(`
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
