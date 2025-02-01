package schema

import "github.com/AwesomeXjs/iq-progress/internal/model"

// OperationSuccessSchema defines the structure of a response for a successful operation.
type OperationSuccessSchema struct {
	Title   string `json:"title" example:"status"`
	Data    int    `json:"data" example:"1000"`
	Request string `json:"request" example:"/api/v1/send"`
	Time    string `json:"time" example:"2023-01-01 00:00:00"`
}

// GetOperationsSchema defines the structure for the response containing a list of operations.
type GetOperationsSchema struct {
	Title   string            `json:"title" example:"status"`
	Data    []model.Operation `json:"data"`
	Request string            `json:"request" example:"/api/v1/send"`
	Time    string            `json:"time" example:"2023-01-01 00:00:00"`
}
