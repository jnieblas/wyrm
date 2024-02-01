package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type Script struct {
	Name        string
	Path        string
	Command     string
	Description string
}

func CreateScript(name string, path string, command string, description string) Script {
	return Script{
		Name:        name,
		Path:        path,
		Command:     command,
		Description: description,
	}
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
	jsonData, err := json.Marshal(script)
	if err != nil {
		log.Fatal("Error parsing script:", err)
	}

	return string(jsonData)
}
