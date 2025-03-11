package config

import (
	"github.com/saipulmuiz/mpio-test/pkg/serror"
	"github.com/saipulmuiz/mpio-test/service/handler/rest"
	"github.com/saipulmuiz/mpio-test/service/repository"
	"github.com/saipulmuiz/mpio-test/service/usecase"
)

func (cfg *Config) InitService() (errx serror.SError) {
	userRepo := repository.NewUserRepo(cfg.DB)
	transactionRepo := repository.NewTransactionRepo(cfg.DB)
	transactionUsecase := usecase.NewTransactionUsecase(cfg.DB, userRepo, transactionRepo)

	route := rest.CreateHandler(
		transactionUsecase,
	)

	cfg.Server = route

	return nil
}
