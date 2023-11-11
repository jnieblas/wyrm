package dao

import (
	"fmt"
	"log"

	"github.com/jnieblas/wyrm/driver"
	"github.com/jnieblas/wyrm/dto"
)

func CreateScript(script *dto.Script) {
	db := driver.GetConnection()
	defer db.Close()

	sql := "INSERT INTO scripts (name, path, command, description) VALUES (?, ?, ?, ?)"
	// Insert user data into the database
	_, err := db.Exec(sql, script.Name, script.Path, script.Command, script.Description)

	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", sql)
	}

	fmt.Println("Script created successfully.")

	GetScripts()
}

func GetScripts() []dto.Script {
	db := driver.GetConnection()
	defer db.Close()

	sql := "SELECT * FROM scripts"
	rows, err := db.Query(sql)

	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to execute SQL: %s", sql)
	}
	defer rows.Close()
	log.Println("Scripts fetched successfully.")

	var scripts []dto.Script

	for rows.Next() {
		script := dto.Script{}
		script.MapRow(rows)
		scripts = append(scripts, script)
	}

	return scripts
}

// func GetScript(name string) dto.Script {

// }
