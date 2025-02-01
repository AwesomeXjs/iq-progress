package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Test(ctx echo.Context) error {
	fmt.Println("test")
	return ctx.JSON(200, "test")
}
