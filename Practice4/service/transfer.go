package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func TransferBalance(db *sqlx.DB, fromID int, toID int, amount float64) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// Проверка баланса отправителя
	var fromBalance float64
	err = tx.Get(&fromBalance, "SELECT balance FROM users WHERE id = $1 FOR UPDATE", fromID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("sender not found")
	}

	if fromBalance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient funds")
	}

	// Проверка получателя
	var toExists bool
	err = tx.Get(&toExists, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", toID)
	if err != nil || !toExists {
		tx.Rollback()
		return fmt.Errorf("receiver not found")
	}

	// Списываем и добавляем
	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
