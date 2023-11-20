package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
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
func PrintTableMetadata(tableName string) {
	db := GetConnection()

	sql := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
