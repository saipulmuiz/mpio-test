package models

import "time"

type Transaction struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"index"`
	Amount    float64   `gorm:"not null"`
	Type      string    `gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WithdrawRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type GetBalanceRequest struct {
	UserID int64 `json:"user_id"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}
