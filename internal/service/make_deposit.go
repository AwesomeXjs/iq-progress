package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/converter"
	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

// MakeDeposit handles adding funds to a user's balance.
func (s *Service) MakeDeposit(ctx context.Context, request model.DepositRequest) (int, error) {

	const mark = "Service.MakeDeposit"

	var balance int
	err := s.TxManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		balance, errTx = s.Repo.AddToBalance(ctx, request.UserID, request.Amount)
		if errTx != nil {
			logger.Error("failed to add to balance", mark, zap.Error(errTx))
			return errTx
		}

		errTx = s.Repo.AddTransaction(ctx,
			converter.ToTxData(
				request.UserID,
				request.UserID,
				request.Amount),
			"deposit")
		if errTx != nil {
			logger.Error("failed to add transaction", mark, zap.Error(errTx))
			return errTx
		}

		return nil
	})
	if err != nil {
		logger.Error("failed to execute transaction", mark, zap.Error(err))
		return 0, err
	}

	return balance, nil
}
