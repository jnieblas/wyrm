package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func getConnection() *sql.DB {
	dbPath, err := getDatabaseFilePath()
	if err != nil {
		fmt.Println("Unable to provision wyrm DB.")
		log.Fatal("Encountered error while getting database file path:", err)
	}

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Println(err)
		log.Fatal("Could not open database")
	}

	return db
}

// For debugging purposes only
func printTableMetadata(tableName string) {
	db := getConnection()
	defer CloseDb(db)
	pragmaSql := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(pragmaSql)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseCursor(rows)

	for rows.Next() {
		fmt.Println(rows)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func getDatabaseFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dbPath := filepath.Join(homeDir, ".wyrm_db")

	return dbPath, nil
}
