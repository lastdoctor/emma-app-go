package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func (m UserModel) Insert(user *User) error {
	query := `
INSERT INTO users (first_name, last_name)
VALUES ($1, $2)
RETURNING id`
	args := []interface{}{user.FirstName, user.LastName}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM users WHERE id = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	log.Println(id)
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
