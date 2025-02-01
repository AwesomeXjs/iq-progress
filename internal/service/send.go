package service

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/converter"
	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

func (s *Service) Send(ctx context.Context, request model.SendRequest) (int, error) {

	const mark = "Service.Send"

	var balance int
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		_, errTx = s.repo.AddToBalance(ctx, request.Receiver, request.Amount)
		if errTx != nil {
			logger.Error("failed to add to balance", mark, zap.Error(errTx))
			return errTx
		}

		balance, errTx = s.repo.RemoveFromBalance(ctx, request.Sender, request.Amount)
		if errTx != nil {
			logger.Error("failed to remove from balance", mark, zap.Error(errTx))
			return errTx
		}

		errTx = s.repo.AddTransaction(ctx, converter.ToTxData(
			request.Sender,
			request.Receiver,
			request.Amount),
			"send")
		if errTx != nil {
			logger.Error("failed to add transaction", mark, zap.Error(errTx))
			return errTx
		}

		errTx = s.repo.AddTransaction(ctx, converter.ToTxData(
			request.Sender,
			request.Receiver,
			request.Amount),
			"receive")
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
