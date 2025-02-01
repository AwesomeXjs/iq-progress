package handler

import (
	"net/http"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) MakeDeposit(ctx echo.Context) error {

	const mark = "Handler.MakeDeposit"

	var Request model.DepositRequest
	if err := ctx.Bind(&Request); err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to bind request", nil)
	}

	_, err := govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Error("failed to validate request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusUnprocessableEntity, "failed to validate request", err.Error())
	}

	balance, err := h.svc.MakeDeposit(ctx.Request().Context(), Request)
	if err != nil {
		return ErrorValidation(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "success", balance)
}
