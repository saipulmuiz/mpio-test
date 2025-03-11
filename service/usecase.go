package service

import (
	"github.com/saipulmuiz/mpio-test/models"
	"github.com/saipulmuiz/mpio-test/pkg/serror"
)

type TransactionUsecase interface {
	GetBalance(request *models.GetBalanceRequest) (transaction *models.GetBalanceResponse, errx serror.SError)
	Withdraw(request *models.WithdrawRequest) (errx serror.SError)
}
