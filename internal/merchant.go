package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Merchant struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	IconUrl     string `json:"icon_url"`
	FunnyGifUrl string `json:"funny_gif_url"`
}

type MerchantStorage interface {
	Create(displayName string, iconUrl string, funnyGifUrl string) (*Merchant, error)
}

type Merchants interface {
	Create(merchant *Merchant) (*Merchant, error)
	GetById(id int64) (*Merchant, error)
}

type MerchantRepository struct {
	DB      *sql.DB
	Storage MerchantStorage
}

func (m MerchantRepository) Create(merchant *Merchant) (*Merchant, error) {
	createdMerchant, err := m.Storage.Create(merchant.DisplayName, merchant.IconUrl, merchant.FunnyGifUrl)
	if err != nil {
		return nil, fmt.Errorf("can't create merchant in storage: %w", err)
	}

	return createdMerchant, nil
}

func (m MerchantRepository) GetById(id int64) (*Merchant, error) {
	if id < 1 {
		// return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM merchant where id = $1`
	var merchant Merchant
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&merchant.ID, &merchant.DisplayName, &merchant.FunnyGifUrl,
		&merchant.FunnyGifUrl)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &merchant, nil
}
