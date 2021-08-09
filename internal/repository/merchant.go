package repository

import (
	"github.com/google/uuid"
	"time"
)

// Merchant is a JSON merchant
type Merchant struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName" validate:"required"`
	IconUrl     string    `json:"iconUrl" validate:"required"`
	FunnyGifUrl string    `json:"funnyGifUrl" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

// DBMerchant is a Postgres merchant
type DBMerchant struct {
	tableName   struct{}  `pg:"merchants" gorm:"primaryKey"`
	ID          uuid.UUID `pg:"id,notnull,pk"`
	DisplayName string    `pg:"display_name"`
	IconUrl     string    `pg:"icon_url"`
	FunnyGifUrl string    `pg:"funny_gif_url"`
	CreatedAt   time.Time `pg:"created_at,notnull"`
}

// ToDB converts Merchent to DBMerchant
func (merchant *Merchant) ToDB() *DBMerchant {
	return &DBMerchant{
		ID:          merchant.ID,
		DisplayName: merchant.DisplayName,
		IconUrl:     merchant.IconUrl,
		FunnyGifUrl: merchant.FunnyGifUrl,
		CreatedAt:   merchant.CreatedAt,
	}
}

// TableName overrides default table name for gorm
func (DBMerchant) TableName() string {
	return "merchants"
}

// ToWeb converts DBUser to User
func (dbMerchant *Merchant) ToWeb() *Merchant {
	return &Merchant{
		ID:          dbMerchant.ID,
		DisplayName: dbMerchant.DisplayName,
		IconUrl:     dbMerchant.IconUrl,
		FunnyGifUrl: dbMerchant.FunnyGifUrl,
		CreatedAt:   dbMerchant.CreatedAt,
	}
}
