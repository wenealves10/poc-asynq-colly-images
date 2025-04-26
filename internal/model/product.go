package model

import "time"

type Product struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	Platform     string        `json:"platform"`
	ExternalID   string        `json:"external_id"`
	ExternalLink string        `json:"external_link"`
	Images       []string      `json:"images"`
	IsActive     bool          `json:"is_active"`
	CurrentPrice float64       `json:"current_price"`
	Reviews      int           `json:"reviews"`
	Status       ProductStatus `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
}
