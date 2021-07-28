package domain

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	MerchantId  int64     `json:"merchant_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
