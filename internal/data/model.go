package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Model struct {
	Merchant    MerchantModel
	User        UserModel
	Transaction TransactionModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		User:        UserModel{DB: db},
		Merchant:    MerchantModel{DB: db},
		Transaction: TransactionModel{DB: db},
	}
}
