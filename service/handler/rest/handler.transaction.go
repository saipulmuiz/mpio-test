package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/mpio-test/models"
	"github.com/saipulmuiz/mpio-test/pkg/serror"
	"github.com/saipulmuiz/mpio-test/pkg/utils/utint"
	"github.com/saipulmuiz/mpio-test/service/helper"
)

func (h *Handler) GetBalance(ctx *gin.Context) {
	var (
		errx serror.SError
	)

	userId := utint.StringToInt(ctx.Param("userId"), 0)

	req := models.GetBalanceRequest{
		UserID: userId,
	}

	data, errx := h.transactionUsecase.GetBalance(&req)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Get balance successfully",
		Data:    data,
	})
}

func (h *Handler) Withdraw(ctx *gin.Context) {
	var (
		request models.WithdrawRequest
		errx    serror.SError
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errx = serror.NewFromErrori(http.StatusBadRequest, err)
		errx.AddComments("[handler][Withdraw] while BodyJSONBind")
		handleError(ctx, errx.Code(), errx)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		validationMessages := helper.BuildAndGetValidationMessage(err)
		handleValidationError(ctx, validationMessages)

		return
	}

	errx = h.transactionUsecase.Withdraw(&request)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Message: "Withdraw request successfully",
	})
}
