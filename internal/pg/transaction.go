package pg

import (
	"github.com/google/uuid"
	"time"
)

// TransactionRepo is store for transactions
//go:generate mockery --dir . --name FileContentRepo --output ./mocks
type TransactionRepo interface {
	GetTransaction(context.Context, uuid.UUID) (*pg.Transaction, error)
	CreateTransaction(context.Context, uuid.UUID) (*pg.Transaction, error)
	UpdateTransaction(context.Context, uuid.UUID) (*pg.Transaction, error)
	DeleteMerchant(context.Context, uuid.UUID) error
}

// Transaction is a JSON transaction
type Transaction struct {
	ID          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	MerchantId  uuid.UUID `json:"merchantId"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// DBTransaction is a Postgres transaction
type DBTransaction struct {
	tableName   struct{}  `pg:"transactions" gorm:"primaryKey"`
	ID          uuid.UUID `pg:"id"`
	UserId      uuid.UUID `pg:"user_id"`
	MerchantId  uuid.UUID `pg:"merchant_id"`
	Amount      int       `pg:"amount"`
	Description string    `pg:"description"`
	CreatedAt   time.Time `pg:"created_at"`
}

// ToDB converts Transaction to DBTransaction
func (transaction *Transaction) ToDB() *DBTransaction {
	return &DBTransaction{
		ID:          transaction.ID,
		UserId:      transaction.UserId,
		MerchantId:  transaction.MerchantId,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		CreatedAt:   transaction.CreatedAt,
	}
}

// TableName overrides default table name for gorm
func (DBTransaction) TableName() string {
	return "transactions"
}

// ToWeb converts DBUser to User
func (dbTransaction *DBTransaction) ToWeb() *Transaction {
	return &Transaction{
		ID:          dbTransaction.ID,
		UserId:      dbTransaction.UserId,
		MerchantId:  dbTransaction.MerchantId,
		Amount:      dbTransaction.Amount,
		Description: dbTransaction.Description,
		CreatedAt:   dbTransaction.CreatedAt,
	}
}
