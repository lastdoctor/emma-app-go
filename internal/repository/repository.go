package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/lastdoctor/emma-app-go/internal/repository/pg"
)

type Store struct {
	Pg *pg.DB
}

//func New(ctx context.Context) (*Store, error) {
//	// Connect to Postgres
//	pgDB, err := pg.Dial()
//	if err != nil {
//		return nil, err
//	}
//
//	// Run Postgres migrations
//	if  {
//
//	}
//}

// UserRepo is repository for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	GetUser(context.Context, uuid.UUID) (*DBUser, error)
	CreateUser(context.Context, *DBUser) (*DBUser, error)
	UpdateUser(context.Context, *DBUser) (*DBUser, error)
	DeleteUser(context.Context, uuid.UUID) error
}

// MerchantRepo is repository for merchants
//go:generate mockery --dir . --name MerchantRepo --output ./mocks
type MerchantRepo interface {
	GetMerchant(context.Context, uuid.UUID) (*Merchant, error)
	CreateMerchant(context.Context, *Merchant) (*Merchant, error)
	UpdateMerchant(context.Context, *Merchant) (*Merchant, error)
	DeleteMerchant(context.Context, uuid.UUID) error
}

// TransactionRepo is repository for transactions
//go:generate mockery --dir . --name FileContentRepo --output ./mocks
type TransactionRepo interface {
	GetTransaction(context.Context, uuid.UUID) (*Transaction, error)
	CreateTransaction(context.Context, uuid.UUID) (*Transaction, error)
	UpdateTransaction(context.Context, uuid.UUID) (*Transaction, error)
	DeleteMerchant(context.Context, uuid.UUID) error
}
