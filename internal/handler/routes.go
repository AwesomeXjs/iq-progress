package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	version1 = "/v1"
	baseBath = "/api"
)

// InitRoutes initializes all the routes for the application.
func (h *Handler) InitRoutes(server *echo.Echo) {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group(baseBath)
	{
		v1 := api.Group(version1)
		{
			v1.GET("/test", h.Test)
			v1.POST("/send", h.Send)
			v1.GET("/operations/:id", h.GetLastOperations)
			v1.POST("/deposit", h.MakeDeposit)
		}
	}
}
