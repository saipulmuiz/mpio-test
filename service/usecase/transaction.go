package usecase

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/mpio-test/models"
	"github.com/saipulmuiz/mpio-test/pkg/serror"
	api "github.com/saipulmuiz/mpio-test/service"

	"gorm.io/gorm"
)

type TransactionUsecase struct {
	db              *gorm.DB
	userRepo        api.UserRepository
	transactionRepo api.TransactionRepository
}

func NewTransactionUsecase(
	db *gorm.DB,
	userRepo api.UserRepository,
	transactionRepo api.TransactionRepository,
) api.TransactionUsecase {
	return &TransactionUsecase{
		db:              db,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *TransactionUsecase) GetBalance(request *models.GetBalanceRequest) (transaction *models.GetBalanceResponse, errx serror.SError) {
	balance, err := u.userRepo.GetUserBalance(request.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][GetBalance] Failed to GetUserBalance, [userID: %s]", request.UserID)
		return
	}

	transaction = &models.GetBalanceResponse{
		Balance: balance.Balance,
	}

	return
}

func (u *TransactionUsecase) Withdraw(request *models.WithdrawRequest) (errx serror.SError) {
	// Start transaction
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userBalance, err := u.userRepo.GetUserBalance(request.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][Withdraw] Failed to GetUserBalance, [userID: %s]", request.UserID)
		return
	}

	if err == gorm.ErrRecordNotFound {
		errx = serror.Newi(http.StatusNotFound, "User not found")
		return
	}

	if userBalance.Balance < request.Amount {
		errx = serror.Newi(http.StatusBadRequest, "Insufficient balance")
		return
	}

	_, err = u.userRepo.UpdateUser(tx, request.UserID, &models.User{
		Balance: userBalance.Balance - request.Amount,
	})
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][Withdraw] Failed to UpdateUser, [userID: %s]", request.UserID)
		return
	}

	transactionArgs := &models.Transaction{
		UserID:    request.UserID,
		Amount:    request.Amount,
		Type:      models.TransactionTypeWithdraw,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = u.transactionRepo.CreateTransaction(tx, transactionArgs)
	if err != nil && err != gorm.ErrRecordNotFound {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][Withdraw] Failed to CreateTransaction, [userID: %s]", request.UserID)
		return
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		errx = serror.Newi(http.StatusInternalServerError, "Failed to commit transaction")
		errx.AddComments("[usecase][Withdraw] Failed to commit transaction")
		return
	}

	return
}
