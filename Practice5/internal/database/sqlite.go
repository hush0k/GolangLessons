package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		price REAL NOT NULL,
		genre TEXT NOT NULL
	);`

	if _, err := db.Exec(createTable); err != nil {
		return nil, fmt.Errorf("ошибка создания таблицы: %w", err)
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)

	if count == 0 {
		insertData := `
		INSERT INTO books (title, price, genre) VALUES
		('War and Peace', 500, 'fiction'),
		('1984', 350, 'fiction'),
		('The Hobbit', 400, 'fantasy'),
		('Clean Code', 600, 'programming'),
		('Design Patterns', 700, 'programming'),
		('Harry Potter', 450, 'fantasy'),
		('The Great Gatsby', 300, 'fiction'),
		('Lord of the Rings', 550, 'fantasy');`

		if _, err := db.Exec(insertData); err != nil {
			return nil, fmt.Errorf("ошибка вставки данных: %w", err)
		}
	}

	return db, nil
}
