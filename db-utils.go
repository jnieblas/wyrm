package main

import "database/sql"

func CloseDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic("Could not close DB: " + err.Error())
	}
}

func CloseCursor(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		panic("Could not close DB cursor: " + err.Error())
	}
}
