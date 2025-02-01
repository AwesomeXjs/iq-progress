package schema

import "github.com/AwesomeXjs/iq-progress/internal/model"

type OperationSuccessSchema struct {
	Title   string `json:"title" example:"status"`
	Data    int    `json:"data" example:"1000"`
	Request string `json:"request" example:"/api/v1/send"`
	Time    string `json:"time" example:"2023-01-01 00:00:00"`
}

type GetOperationsSchema struct {
	Title   string            `json:"title" example:"status"`
	Data    []model.Operation `json:"data"`
	Request string            `json:"request" example:"/api/v1/send"`
	Time    string            `json:"time" example:"2023-01-01 00:00:00"`
}
