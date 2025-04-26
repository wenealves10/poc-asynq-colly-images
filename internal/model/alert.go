package model

import "time"

type Alert struct {
	ID           string     `json:"id"`
	ProductID    string     `json:"product_id"`
	UserID       string     `json:"user_id"`
	TargetPrice  float64    `json:"target_price"`
	IsActive     bool       `json:"is_active"`
	Notified     bool       `json:"notified"`
	Clicked      bool       `json:"clicked"`
	SendEmail    bool       `json:"send_email"`
	SendWhatsApp bool       `json:"send_whatsapp"`
	SentAt       *time.Time `json:"sent_at"`
	ClickedAt    *time.Time `json:"clicked_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
