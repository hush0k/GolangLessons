package handlers

import (
	"Practice5/internal/models"
	"Practice5/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BookHandler struct {
	repo *repository.BookRepository
}

func NewBookHandler(repo *repository.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	genre := query.Get("genre")
	sort := query.Get("sort")
	limit := 10
	if limitStr := query.Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}
	offset := 0
	if offsetStr := query.Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}
	filter := models.BookFilter{
		Genre:  genre,
		Sort:   sort,
		Limit:  limit,
		Offset: offset,
	}

	// Получаем книги из БД
	books, err := h.repo.GetBooks(filter)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения данных: %v", err), http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки
	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
		return
	}
}
