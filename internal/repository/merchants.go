package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lastdoctor/emma-app-go/internal/domain"
	"time"
)

type MerchantModel struct {
	DB *sql.DB
}

func (m MerchantModel) Create(merchant *domain.Merchants) error {
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

func (m MerchantModel) GetById(id int64) (*domain.Merchants, error) {
	if id < 1 {
		//return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM merchant where id = $1`
	var merchant domain.Merchants
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&merchant.ID, &merchant.DisplayName, &merchant.FunnyGifUrl,
		&merchant.FunnyGifUrl)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			//return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &merchant, nil
}
