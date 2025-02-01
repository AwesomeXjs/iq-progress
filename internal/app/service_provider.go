package app

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/config"
	"github.com/AwesomeXjs/iq-progress/internal/handler"
	"github.com/AwesomeXjs/iq-progress/internal/repository"
	"github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/pkg/closer"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient/pg"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient/transaction"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

// ServiceProvider manages the initialization and lifecycle of various components
// in the application, including the HTTP config, Postgres config, DB client, transaction manager,
// repository, service, and handler.
type ServiceProvider struct {
	httpConfig *config.HTTPConfig
	pgConfig   *config.PgConfig

	dbClient  dbclient.Client
	txManager dbclient.TxManager

	handler    *handler.Handler
	service    service.IService
	repository repository.IRepository
}

// NewServiceProvider creates a new instance of ServiceProvider.
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

// HTTPConfig returns the HTTP configuration, initializing it if necessary.
func (s *ServiceProvider) HTTPConfig() *config.HTTPConfig {
	const mark = "App.ServiceProvider.HTTPConfig"

	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", mark, zap.Error(err))
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

// PGConfig initializes and returns the PostgresSQL configuration if not already set.
func (s *ServiceProvider) PGConfig() *config.PgConfig {
	const mark = "App.ServiceProvider.PGConfig"

	if s.pgConfig == nil {
		cfg, err := config.NewPgConfig()
		if err != nil {
			logger.Fatal("failed to get pg config", mark, zap.Error(err))
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

// DBClient initializes and returns the database client if not already created.
// It also pings the database to ensure the connection is valid.
func (s *ServiceProvider) DBClient(ctx context.Context) dbclient.Client {
	const mark = "App.ServiceProvider.DBClient"

	if s.dbClient == nil {
		cfg := s.PGConfig()
		dbc, err := pg.New(ctx, cfg.GetDSN())
		if err != nil {
			logger.Fatal("failed to get db client", mark, zap.Error(err))
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping db", mark, zap.Error(err))
		}

		closer.Add(dbc.Close) // Ensures the database client is closed on shutdown
		s.dbClient = dbc
	}
	return s.dbClient
}

// InitTxManager initializes and returns the transaction manager if not already created.
func (s *ServiceProvider) InitTxManager(ctx context.Context) dbclient.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

// InitRepository initializes and returns the repository if not already created.
func (s *ServiceProvider) InitRepository(ctx context.Context) repository.IRepository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx))
	}
	return s.repository
}

// InitService initializes and returns the service if not already created.
func (s *ServiceProvider) InitService(ctx context.Context) service.IService {
	if s.service == nil {
		s.service = service.New(s.InitRepository(ctx), s.InitTxManager(ctx))
	}
	return s.service
}

// InitHandler initializes and returns the handler if not already created.
func (s *ServiceProvider) InitHandler(ctx context.Context) *handler.Handler {
	if s.handler == nil {
		s.handler = handler.New(s.InitService(ctx))
	}
	return s.handler
}
