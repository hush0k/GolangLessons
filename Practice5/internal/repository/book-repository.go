package repository

import (
	"Practice5/internal/models"
	"database/sql"
	"fmt"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBooks(filter models.BookFilter) ([]models.Book, error) {
	query := "SELECT id, title, price, genre FROM books WHERE 1=1"
	args := []interface{}{}
	if filter.Genre != "" {
		query += " AND genre = ?"
		args = append(args, filter.Genre)
	}
	switch filter.Sort {
	case "price_asc":
		query += " ORDER BY price ASC"
	case "price_desc":
		query += " ORDER BY price DESC"
	default:
		query += " ORDER BY id ASC"
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Price, &book.Genre); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации: %w", err)
	}

	return books, nil
}
