package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/AwesomeXjs/iq-progress/internal/middlewares"
	"github.com/AwesomeXjs/iq-progress/pkg/closer"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	// EnvPath is the path to the .env file that contains environment variables.
	EnvPath = ".env.example"
)

// logLevel is a command-line flag for specifying the log level.
var logLevel = flag.String("l", "info", "log level")

type App struct {
	server          *echo.Echo
	serviceProvider *ServiceProvider
}

func New(ctx context.Context) *App {
	const mark = "App.App.New"

	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		// Fatal log in case of failure during dependency initialization
		logger.Fatal("failed to init deps", mark, zap.Error(err))
	}
	return app
}

func (app *App) InitDeps(ctx context.Context) error {
	const mark = "App.App.InitDeps"

	inits := []func(ctx context.Context) error{
		app.InitConfig,
		app.initServiceProvider,
		app.InitEchoServer,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			logger.Error("failed to init deps", mark, zap.Error(err))
			return err
		}
	}

	app.InitRoutes(ctx, app.server)
	return nil
}

func (app *App) Run() error {
	const mark = "App.App.Run"

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := app.runHTTPServer() // Run the HTTP server
	if err != nil {
		logger.Fatal("failed to run http server", mark, zap.Error(err))
	}
	return nil
}

func (app *App) InitConfig(_ context.Context) error {
	const mark = "App.App.InitConfig"

	err := godotenv.Load(EnvPath)
	if err != nil {
		logger.Error("Error loading .env file", mark, zap.String("path", EnvPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return err
}

// InitEchoServer sets up the Echo server and its middleware.
func (app *App) InitEchoServer(_ context.Context) error {
	flag.Parse()                                                 // Parse command-line flags
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel))) // Initialize logger with the specified log level

	app.server = echo.New()              // Create a new Echo server
	app.server.Use(middleware.Recover()) // Middleware for recovering from panics
	app.server.Use(middlewares.Logger)   // Custom logging middleware

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = NewServiceProvider()
	return nil
}

func (app *App) runHTTPServer() error {
	const mark = "App.App.runHTTPServer"

	logger.Info("server listening at %v", mark, zap.String("start", app.serviceProvider.HTTPConfig().Address())) // Log the server address
	return app.server.Start(app.serviceProvider.HTTPConfig().Address())                                          // Start the server at the configured address
}

// InitRoutes sets up the application routes.
func (app *App) InitRoutes(ctx context.Context, server *echo.Echo) {
	app.serviceProvider.InitHandler(ctx).InitRoutes(server) // Initialize routes using the controller
}
