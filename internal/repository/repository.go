package repository

import (
	"database/sql"
	"errors"
	"github.com/lastdoctor/emma-app-go/internal/domain"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Users interface {
	Create(user *Users) error
	GetById(id int64) (*domain.User, error)
}

type UsersRepos struct {
	DB *sql.DB
}

type Merchants interface {
	Create(merchant *domain.Merchants) error
	GetById(id int64) (*domain.Merchants, error)
}

type MerchantsRepos struct {
	DB *sql.DB
}

type Transactions interface {
	Create()
}

type Repositories struct {
	Users        Users
	Merchants    Merchants
	Transactions Transactions
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:        UsersRepos{DB: db},
		Merchants:    MerchantsRepos{DB: db},
		Transactions: nil,
	}
}
