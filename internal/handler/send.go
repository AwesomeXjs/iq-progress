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

// Send - Send
// @Summary Send
// @Tags User
// @Description send money to another user
// @ID Send
// @Accept  json
// @Produce  json
// @Param input body model.SendRequest true "data"
// @Success 200 {object} schema.OperationSuccessSchema
// @Failure 400 {object} utils.Body
// @Failure 422 {object} utils.Body
// @Failure 500 {object} utils.Body
// @Router /api/v1/send [post]
func (h *Handler) Send(ctx echo.Context) error {

	const mark = "Handler.Send"

	var Request model.SendRequest

	if err := ctx.Bind(&Request); err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to bind request", nil)
	}

	_, err := govalidator.ValidateStruct(Request)
	if err != nil {

		logger.Error("failed to validate request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusUnprocessableEntity, "failed to validate request", err.Error())
	}

	balance, err := h.svc.Send(ctx.Request().Context(), Request)
	if err != nil {
		return ErrorValidation(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "success", balance)
}
