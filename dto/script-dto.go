package dto

import (
	"database/sql"
	"log"
)

type Script struct {
	Name        string
	Path        string
	Command     string
	Description string
}

func (script *Script) MapRow(rows *sql.Rows) {
	err := rows.Scan(&script.Name, &script.Path, &script.Command, &script.Description)
	if err != nil {
		log.Println(err)
		log.Fatal("Unable to read results from scripts table.")
	}
	log.Printf("Script Mapping - Name: %s, Path: %s, Command: %s, Description: %s\n", script.Name, script.Path, script.Command, script.Description)
}
