package driver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "testdb.db")

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
