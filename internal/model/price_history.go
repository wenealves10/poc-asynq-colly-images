package model

import "time"

type PriceHistory struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Price     float64   `json:"price"`
	CheckedAt time.Time `json:"checked_at"`
	CreatedAt time.Time `json:"created_at"`
}
