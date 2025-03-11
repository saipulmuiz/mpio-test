package service

import (
	"github.com/saipulmuiz/mpio-test/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserBalance(userID int64) (*models.User, error)
	UpdateUser(tx *gorm.DB, userID int64, user *models.User) (*models.User, error)
}

type TransactionRepository interface {
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error)
}
