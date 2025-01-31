package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	version1 = "v1"
	baseBath = "/api"
)

// InitRoutes initializes all the routes for the application.
func (c *Handler) InitRoutes(server *echo.Echo) {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group(baseBath)
	{
		_ = api.Group(version1)
		{

		}
	}
}
