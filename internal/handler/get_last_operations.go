package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) GetLastOperations(ctx echo.Context) error {
	fmt.Println("YO")

	const mark = "Handler.GetLastOperations"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to bind request", nil)
	}

	var operations []model.Operations

	operations, err = h.svc.GetOperations(ctx.Request().Context(), id)
	if err != nil {
		return utils.Response(ctx, http.StatusInternalServerError, "failed to get operations", nil)
	}

	return utils.Response(ctx, http.StatusOK, "success", operations)
}
