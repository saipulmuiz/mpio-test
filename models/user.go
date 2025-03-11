package models

import (
	"time"
)

type User struct {
	UserID    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	Name      string    `gorm:"not null;size:256" json:"name"`
	Balance   float64   `gorm:"not null;size:256" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBalance struct {
	UserID  int64   `json:"user_id"`
	Balance float64 `json:"balance"`
}
