package main

import (
	"fmt"
	"log"

	"Practice4/config"
	"Practice4/models"
	"Practice4/repository"
	"Practice4/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	// Добавляем пользователей
	alice := models.User{Name: "Alice", Email: "alice@mail.com", Balance: 1000}
	bob := models.User{Name: "Bob", Email: "bob@mail.com", Balance: 500}

	_ = repository.InsertUser(db, alice)
	_ = repository.InsertUser(db, bob)

	users, _ := repository.GetAllUsers(db)
	fmt.Println("Users:", users)

	// Перевод
	err := service.TransferBalance(db, 1, 2, 500)
	if err != nil {
		log.Println("❌ Transaction failed:", err)
	} else {
		log.Println("✅ Transfer successful!")
	}
}
