package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
)

func (s *Service) GetOperations(ctx context.Context, userID int) ([]model.Operation, error) {
	return s.repo.GetOperations(ctx, userID)
}
