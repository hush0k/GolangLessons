package repository

import (
	"Practice4/models"

	"github.com/jmoiron/sqlx"
)

func InsertUser(db *sqlx.DB, user models.User) error {
	query := `
		INSERT INTO users (name, email, balance)
		VALUES (:name, :email, :balance)
	`
	_, err := db.NamedExec(query, user)
	return err
}

func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
	var users []models.User
	err := db.Select(&users, "SELECT * FROM users ORDER BY id")
	return users, err
}

func GetUserByID(db *sqlx.DB, id int) (models.User, error) {
	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}
