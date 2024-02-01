package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Script struct {
	Name        string
	Path        string
	Command     string
	Description string
}

func (script *Script) MapRows(rows *sql.Rows) {
	err := rows.Scan(&script.Name, &script.Path, &script.Command, &script.Description)
	if err != nil {
		fmt.Printf("Could not get script with values:\n%v", script)
		log.Println(err)
		log.Fatal("Unable to read results from scripts table.")
	}
}

func (script *Script) MapRow(row *sql.Row) {
	err := row.Scan(&script.Name, &script.Path, &script.Command, &script.Description)
	if err != nil {
		fmt.Printf("Could not get script with values:%v", script)
		log.Println(err)
		log.Fatal("Unable to read results from scripts table.")
	}
}

func (script *Script) String() string {
	return fmt.Sprintf(`
Name: '%s'
Path: '%s'
Command: '%s'
Description: '%s'`, script.Name, script.Path, script.Command, script.Description)
}
