package repository

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
)

// IRepository defines the interface for repository operations related to balance and transactions.
type IRepository interface {
	// AddToBalance adds the specified amount to the user's balance.
	AddToBalance(ctx context.Context, userID int, amount int) (int, error)
	// RemoveFromBalance removes the specified amount from the user's balance.
	RemoveFromBalance(ctx context.Context, userID int, amount int) (int, error)
	// AddTransaction records a new transaction with the specified data and type.
	AddTransaction(ctx context.Context, data *model.TxData, txType string) error
	// GetOperations retrieves the most recent operations for the specified user.
	GetOperations(ctx context.Context, userID int) ([]model.Operation, error)
}

// Repository implements the IRepository interface for interacting with the database.
type Repository struct {
	db dbclient.Client
}

// New creates a new instance of the Repository with the provided database client.
func New(db dbclient.Client) IRepository {
	return &Repository{
		db: db,
	}
}
