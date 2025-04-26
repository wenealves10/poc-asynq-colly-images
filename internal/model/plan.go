package model

import "time"

type Plan struct {
	ID                  string    `json:"id"`
	ExternalID          string    `json:"external_id"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	IsFree              bool      `json:"is_free"`
	MaxProducts         int       `json:"max_products"`
	MaxAlertsPerProduct int       `json:"max_alerts_per_product"`
	Platforms           []string  `json:"platforms"`
	PriceMonthly        float64   `json:"price_monthly"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
