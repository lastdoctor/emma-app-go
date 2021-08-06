package repository

import (
	"context"
	"database/sql"
	"time"
)

type Transaction struct {
	ID          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	MerchantId  int64     `json:"merchant_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (m TransactionModel) Create(transaction *Transaction) error {
	query := `
INSERT INTO user (user_id, merchant_id, amount, description, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	args := []interface{}{transaction.UserId, transaction.MerchantId, transaction.Amount, transaction.Description,
		transaction.Date}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&transaction.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m TransactionModel) GetById(id int64) (*Transaction, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM transaction WHERE id = $1`
	var transaction Transaction
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&transaction.ID, &transaction.UserId, &transaction.MerchantId,
		&transaction.MerchantId, &transaction.Amount, &transaction.Description, &transaction.Date)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &transaction, nil
}
