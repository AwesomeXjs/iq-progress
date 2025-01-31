package main

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/app"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

func main() {

	const mark = "main"

	ctx := context.Background()

	myApp := app.New(ctx)

	err := myApp.Run()
	if err != nil {
		logger.Fatal("failed to run app", mark, zap.Error(err))
	}
}
