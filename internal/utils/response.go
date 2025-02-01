package utils

import (
	"time"

	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Body struct defines the structure of the response body.
// It includes a title, detail about the error or message, the request URI, and the current time.
type Body struct {
	Title   string `json:"title"`
	Data    any    `json:"data"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

// Response is a utility function to send JSON responses.
// It accepts the HTTP context, a status code, a message, and detailed information to send in the response.
func Response(ctx echo.Context, statusCode int, message string, data interface{}) error {
	const mark = "Response Wrapper"
	// Construct the response object using the provided parameters.
	err := ctx.JSON(statusCode, Body{
		Title:   message,
		Data:    data,
		Request: ctx.Request().RequestURI,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	})

	// If there's an error in writing the response, log it and return the error.
	if err != nil {
		logger.Error("failed to write response", mark, zap.String("error", err.Error()))
		return err
	}

	// Return nil indicating that the response was successfully written.
	return nil
}
