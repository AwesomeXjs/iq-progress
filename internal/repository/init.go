package repository

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
)

type IRepository interface {
	AddToBalance(ctx context.Context, userID int, amount int) (int, error)
	RemoveFromBalance(ctx context.Context, userID int, amount int) (int, error)
	AddTransaction(ctx context.Context, data *model.TxData, txType string) error
	GetOperations(ctx context.Context, userID int) ([]model.Operations, error)
}

type Repository struct {
	db dbClient.Client
}

func New(db dbClient.Client) IRepository {
	return &Repository{
		db: db,
	}
}
