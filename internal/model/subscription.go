package model

import "time"

type Subscription struct {
	ID         string             `json:"id"`
	ExternalID string             `json:"external_id"`
	UserID     string             `json:"user_id"`
	PlanID     string             `json:"plan_id"`
	IsActive   bool               `json:"is_active"`
	Status     SubscriptionStatus `json:"status"`
	StartedAt  time.Time          `json:"started_at"`
	ExpiresAt  time.Time          `json:"expires_at"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
