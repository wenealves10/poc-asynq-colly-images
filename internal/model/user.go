package model

import "time"

type User struct {
	ID          string     `json:"id"`
	CustomerID  string     `json:"customer_id"`
	Role        Role       `json:"role"`
	IsActive    bool       `json:"is_active"`
	FirebaseUID string     `json:"firebase_uid"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
