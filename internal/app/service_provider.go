package app

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/config"
	"github.com/AwesomeXjs/iq-progress/internal/handler"
	"github.com/AwesomeXjs/iq-progress/internal/repository"
	"github.com/AwesomeXjs/iq-progress/internal/service"
	"github.com/AwesomeXjs/iq-progress/pkg/closer"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient/pg"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient/transaction"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

type ServiceProvider struct {
	httpConfig *config.HTTPConfig
	pgConfig   *config.PgConfig

	dbClient  dbClient.Client
	txManager dbClient.TxManager

	handler    *handler.Handler
	service    service.IService
	repository repository.IRepository
}

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
func (s *ServiceProvider) DBClient(ctx context.Context) dbClient.Client {
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

func (s *ServiceProvider) InitTxManager(ctx context.Context) dbClient.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *ServiceProvider) InitRepository(ctx context.Context) repository.IRepository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx))
	}
	return s.repository
}

func (s *ServiceProvider) InitService(ctx context.Context) service.IService {
	if s.service == nil {
		s.service = service.New(s.InitRepository(ctx), s.InitTxManager(ctx))
	}
	return s.service
}

func (s *ServiceProvider) InitHandler(ctx context.Context) *handler.Handler {
	if s.handler == nil {
		s.handler = handler.New(s.InitService(ctx))
	}
	return s.handler
}
