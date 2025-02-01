package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
)

// GetOperations retrieves the transaction history of a user.
func (s *Service) GetOperations(ctx context.Context, userID int) ([]model.Operation, error) {
	return s.Repo.GetOperations(ctx, userID)
}
