package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/repository"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
)

// IService defines the interface for transaction-related operations.
type IService interface {
	Send(ctx context.Context, request model.SendRequest) (int, error)
	MakeDeposit(ctx context.Context, request model.DepositRequest) (int, error)
	GetOperations(ctx context.Context, userID int) ([]model.Operation, error)
}

// Service implements IService and handles business logic.
type Service struct {
	Repo      repository.IRepository
	TxManager dbclient.TxManager
}

// New creates a new Service instance.
func New(repo repository.IRepository, txManager dbclient.TxManager) IService {
	return &Service{
		Repo:      repo,
		TxManager: txManager,
	}
}
