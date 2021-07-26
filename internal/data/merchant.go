package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Merchant struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	IconUrl     string `json:"icon_url"`
	FunnyGifUrl string `json:"funny_gif_url"`
}

type MerchantModel struct {
	DB *sql.DB
}

type IMerchant interface {
	Insert(merchant *Merchant)
	Get(id int64)
}

func (m MerchantModel) Insert(merchant *Merchant) error {
	query := `
INSERT INTO merchant (display_name, icon_url, funny_gif_url)
VALUES($1, $2)
RETURNING id`
	args := []interface{}{merchant.DisplayName, merchant.IconUrl, merchant.FunnyGifUrl}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&merchant.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m MerchantModel) Get(id int64) (*Merchant, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
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
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &merchant, nil
}
