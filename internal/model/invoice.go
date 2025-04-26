package model

import "time"

type Invoice struct {
	ID             string        `json:"id"`
	ExternalID     string        `json:"external_id"`
	UserID         string        `json:"user_id"`
	SubscriptionID string        `json:"subscription_id"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
	Amount         float64       `json:"amount"`
	Status         InvoiceStatus `json:"status"`
	DueDate        time.Time     `json:"due_date"`
	PaidAt         *time.Time    `json:"paid_at"`
	CreatedAt      time.Time     `json:"created_at"`
}
