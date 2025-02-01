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

// MakeDeposit - MakeDeposit
// @Summary MakeDeposit
// @Tags User
// @Description make deposit for user
// @ID MakeDeposit
// @Accept  json
// @Produce  json
// @Param input body model.DepositRequest true "deposit info"
// @Success 200 {object} schema.OperationSuccessSchema
// @Failure 400 {object} utils.Body
// @Failure 422 {object} utils.Body
// @Failure 500 {object} utils.Body
// @Router /api/v1/deposit [post]
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
