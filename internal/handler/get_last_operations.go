package handler

import (
	"net/http"
	"strconv"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// GetLastOperations - GetLastOperations
// @Summary GetLastOperations
// @Tags User
// @Description get last operations for user
// @ID GetLastOperations
// @Accept  json
// @Produce  json
// @Param id path int false "user id"
// @Success 200 {object} schema.GetOperationsSchema
// @Failure 400 {object} utils.Body
// @Failure 422 {object} utils.Body
// @Failure 500 {object} utils.Body
// @Router /api/v1/operations/{id} [get]
func (h *Handler) GetLastOperations(ctx echo.Context) error {

	const mark = "Handler.GetLastOperations"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to bind request", nil)
	}

	var operations []model.Operation

	operations, err = h.svc.GetOperations(ctx.Request().Context(), id)
	if err != nil {
		return utils.Response(ctx, http.StatusInternalServerError, "failed to get operations", nil)
	}

	return utils.Response(ctx, http.StatusOK, "success", operations)
}
