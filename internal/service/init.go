package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/internal/repository"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
)

type IService interface {
	Send(ctx context.Context, request model.SendRequest) (int, error)
	MakeDeposit(ctx context.Context, request model.DepositRequest) (int, error)
	GetOperations(ctx context.Context, userID int) ([]model.Operation, error)
}

type Service struct {
	repo      repository.IRepository
	txManager dbClient.TxManager
}

func New(repo repository.IRepository, txManager dbClient.TxManager) IService {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}
