package handler

import "github.com/labstack/echo/v4"

type IHandler interface {
	InitRoutes(server *echo.Echo)
}

// Handler handles the authentication and header-related operations.
type Handler struct {
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New() IHandler {
	return &Handler{}
}
