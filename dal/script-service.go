package dal

import (
	"fmt"
	"log"

	"github.com/jnieblas/wyrm/driver"
)

type Script struct {
	Name        string
	Path        string
	Command     string
	Description string
}

func CreateScript() Script {
	db := driver.GetConnection()
	defer db.Close()
	// testing insert
	name := "test script"
	path := "a path"
	command := "a nice command!"
	description := "a desc"

	sql := "INSERT INTO scripts (name, path, command, description) VALUES (?, ?, ?, ?)"
	// Insert user data into the database
	_, err := db.Exec(sql, name, path, command, description)

	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", sql)
	}

	fmt.Println("Script stored successfully.")

	return Script{
		Name:        name,
		Path:        path,
		Command:     command,
		Description: description,
	}
}

func GetScripts() []Script {
	db := driver.GetConnection()
	defer db.Close()

	sql := "SELECT * FROM scripts"
	rows, err := db.Query(sql)

	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", sql)
	}
	defer rows.Close()

	var scripts []Script

	for rows.Next() {
		var script Script
		err := rows.Scan(&script.Name, &script.Path, &script.Command, &script.Description)
		if err != nil {
			log.Println(err)
			log.Fatal("Unable to read results from scripts table.")
		}

		log.Printf("Name: %s, path: %s, command: %s, description: %s\n", script.Name, script.Path, script.Command, script.Description)
	}

	log.Println("Script fetched successfully.")

	return scripts
}
