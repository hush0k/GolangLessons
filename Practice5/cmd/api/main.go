package main

import (
	"Practice5/internal/database"
	"Practice5/internal/handlers"
	"Practice5/internal/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.NewSQLiteDB("books.db")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	log.Println("Подключение к БД успешно!")

	bookRepo := repository.NewBookRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepo)

	router := mux.NewRouter()
	router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")

	port := "8088"
	log.Printf("Сервер запущен на http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
