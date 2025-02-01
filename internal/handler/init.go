package handler

import (
	"net/http"

	"github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Handler handles the authentication and header-related operations.
type Handler struct {
	svc service.IService
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New(svc service.IService) *Handler {
	return &Handler{
		svc: svc,
	}
}

// ErrorValidation handles the error response based on the error type.
func ErrorValidation(ctx echo.Context, err error) error {
	switch {
	case errors.Is(err, utils.ErrNotEnoughBalance):
		return utils.Response(ctx, http.StatusUnprocessableEntity, "not enough balance", nil)
	case errors.Is(err, utils.ErrUserNotFound):
		return utils.Response(ctx, http.StatusNotFound, "user not found", nil)
	case errors.Is(err, utils.ErrSenderNotFound):
		return utils.Response(ctx, http.StatusInternalServerError, "sender not found", nil)
	default:
		return utils.Response(ctx, http.StatusBadRequest, "failed to remove from balance", err.Error())
	}
}
