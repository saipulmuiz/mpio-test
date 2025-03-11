package repository

import (
	"github.com/saipulmuiz/mpio-test/models"
	api "github.com/saipulmuiz/mpio-test/service"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) api.TransactionRepository {
	return &transactionRepo{db}
}

func (u *transactionRepo) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error) {
	var result *gorm.DB
	if tx != nil {
		result = tx.Model(&transaction).Clauses(clause.Returning{}).Create(&transaction)
	} else {
		result = u.db.Model(&transaction).Clauses(clause.Returning{}).Create(&transaction)
	}

	return transaction, result.Error
}
