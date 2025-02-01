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

// LogLevel is a command-line flag for specifying the log level.
var LogLevel = flag.String("l", "info", "log level")

// App represents the main application with its server and service provider.
type App struct {
	server          *echo.Echo       // Echo server instance
	serviceProvider *ServiceProvider // Service provider for dependency management
}

// New initializes a new App instance and its dependencies.
func New(ctx context.Context) *App {
	const mark = "App.App.New"

	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		logger.Fatal("failed to init deps", mark, zap.Error(err))
	}
	return app
}

// InitDeps initializes all the required dependencies for the application.
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

// Run starts the application and runs the HTTP server.
func (app *App) Run() error {
	const mark = "App.App.Run"

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := app.runHTTPServer()
	if err != nil {
		logger.Fatal("failed to run http server", mark, zap.Error(err))
	}
	return nil
}

// InitConfig loads environment variables from the .env file.
func (app *App) InitConfig(_ context.Context) error {
	const mark = "App.App.InitConfig"

	err := godotenv.Load(EnvPath)
	if err != nil {
		logger.Error("Error loading .env file", mark, zap.String("path", EnvPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return err
}

// InitEchoServer sets up the Echo server and middleware.
func (app *App) InitEchoServer(_ context.Context) error {
	flag.Parse()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(LogLevel)))

	app.server = echo.New()
	app.server.Use(middleware.Recover())
	app.server.Use(middlewares.Logger)

	return nil
}

// initServiceProvider initializes the service provider.
func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = NewServiceProvider()
	return nil
}

// runHTTPServer starts the HTTP server on the configured address.
func (app *App) runHTTPServer() error {
	const mark = "App.App.runHTTPServer"

	logger.Info("server listening at %v", mark, zap.String("start", app.serviceProvider.HTTPConfig().Address()))
	return app.server.Start(app.serviceProvider.HTTPConfig().Address())
}

// InitRoutes sets up application routes using the service provider.
func (app *App) InitRoutes(ctx context.Context, server *echo.Echo) {
	app.serviceProvider.InitHandler(ctx).InitRoutes(server)
}
