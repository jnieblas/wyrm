package main

import (
	"fmt"
	"log"
	"os"
)

func getScriptsExecutor() []Script {
	db := getConnection()

	defer CloseDb(db)

	selectSql := "SELECT * FROM scripts"
	rows, err := db.Query(selectSql)

	if err != nil {
		fmt.Println("Unable to get scripts.")
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", selectSql)
	}

	defer CloseCursor(rows)

	log.Println("Scripts fetched successfully.")

	var scripts []Script

	for rows.Next() {
		script := Script{}
		script.MapRows(rows)
		scripts = append(scripts, script)
	}

	return scripts
}

func getScriptExecutor(name string) Script {
	db := getConnection()
	defer CloseDb(db)

	selectSql := "SELECT * FROM scripts WHERE name = ?"
	row := db.QueryRow(selectSql, name)

	log.Printf("Script '%s' fetched successfully.", name)

	script := Script{}
	script.MapRow(row)
	return script
}

func createScriptExecutor(script *Script) {
	db := getConnection()
	defer CloseDb(db)

	insertSql := `
		INSERT INTO scripts (name, path, command, description) 
		VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(insertSql, script.Name, script.Path, script.Command, script.Description)

	if err != nil {
		log.Println(err)
		fmt.Printf("Unable to execute SQL: %s\n", insertSql)
		os.Exit(1)
	}
}

func updateScriptExecutor(script *Script) {
	db := getConnection()
	defer CloseDb(db)

	updateSql := `
		UPDATE scripts
		SET name = ?, path = ?, command = ?, description = ?
		WHERE name = ?
	`
	_, err := db.Exec(updateSql, script.Name, script.Path, script.Command, script.Description, script.Name)

	if err != nil {
		log.Println(err)
		log.Printf("Unable to execute SQL: %s\n", updateSql)
		fmt.Printf("Unable to update script with values:\n%v", script)
		os.Exit(1)
	}
}
