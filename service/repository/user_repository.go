package repository

import (
	"github.com/saipulmuiz/mpio-test/models"
	api "github.com/saipulmuiz/mpio-test/service"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) api.UserRepository {
	return &userRepo{db}
}

func (u *userRepo) GetUserBalance(userID int64) (*models.User, error) {
	var user models.User
	err := u.db.Where("user_id = ?", userID).First(&user).Error
	return &user, err
}

func (u *userRepo) UpdateUser(tx *gorm.DB, userID int64, user *models.User) (*models.User, error) {
	var result *gorm.DB
	if tx != nil {
		result = tx.Model(&models.User{}).Clauses(clause.Returning{}).Where("user_id", userID).Updates(user)
	} else {
		result = u.db.Model(&models.User{}).Clauses(clause.Returning{}).Where("user_id", userID).Updates(user)
	}

	return user, result.Error
}
