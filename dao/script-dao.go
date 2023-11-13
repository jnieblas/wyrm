package dao

import (
	"fmt"
	"log"
	"os"

	"github.com/jnieblas/wyrm/driver"
	"github.com/jnieblas/wyrm/dto"
)

func GetScripts() []dto.Script {
	db := driver.GetConnection()
	defer db.Close()

	sql := "SELECT * FROM scripts"
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("Unable to get scripts.")
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", sql)
	}
	defer rows.Close()
	log.Println("Scripts fetched successfully.")

	var scripts []dto.Script

	for rows.Next() {
		script := dto.Script{}
		script.MapRows(rows)
		scripts = append(scripts, script)
	}

	return scripts
}

func GetScript(name string) dto.Script {
	db := driver.GetConnection()
	defer db.Close()

	sql := "SELECT * FROM scripts WHERE name = '?'"
	fmt.Println("sql:", sql)
	row := db.QueryRow(sql, name)

	log.Printf("Script '%s' fetched successfully.\n", name)

	script := dto.Script{}
	script.MapRow(row)
	return script
}

func CreateScript(script *dto.Script) {
	db := driver.GetConnection()
	defer db.Close()

	sql := `
		INSERT INTO scripts (name, path, command, description) 
		VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(sql, script.Name, script.Path, script.Command, script.Description)

	if err != nil {
		log.Println(err)
		fmt.Printf("Unable to execute SQL: %s\n", sql)
		os.Exit(1)
	}
}

func UpdateScript(script *dto.Script) {
	db := driver.GetConnection()
	defer db.Close()

	sql := `
		UPDATE scripts
		SET name = ?, path = ?, command = ?, description = ?
		WHERE name = ?
	`
	_, err := db.Exec(sql, script.Name, script.Path, script.Command, script.Description, script.Name)

	if err != nil {
		log.Println(err)
		log.Printf("Unable to execute SQL: %s\n", sql)
		fmt.Printf("Unable to update script with values:\n%v", script)
		os.Exit(1)
	}
}
