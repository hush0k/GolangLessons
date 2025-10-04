package main

import (
	"GolangLessons/Practice2/internal/handlers"
	"GolangLessons/Practice2/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрируем роуты с middleware
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.APIKeyMiddleware(handlers.GetUser)(w, r)
		case http.MethodPost:
			middleware.APIKeyMiddleware(handlers.CreateUser)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
