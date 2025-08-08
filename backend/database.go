package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./messages.db")
	if err != nil {
		panic(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		content TEXT NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
}

func insertMessage(name, content string) error {
	_, err := db.Exec("INSERT INTO messages (name, content) VALUES (?, ?)", name, content)
	return err
}

func getAllMessages() ([]Message, error) {
	rows, err := db.Query("SELECT id, name, content FROM messages ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Name, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
