package repository

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) GetOperations(ctx context.Context, userID int) ([]model.Operations, error) {

	const mark = "Repository.GetOperations"

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	queryBuilder := builder.
		Select(
			TxID,
			SenderID,
			SenderUsername,
			ReceiverID,
			ReceiverUsername,
			TxType,
			TxAmount,
			TxDate,
		).
		From(TxTable).
		LeftJoin(U1+" ON t.from_user_id = u1.id").
		LeftJoin(U2+" ON t.to_user_id = u2.id").
		Where("(t.from_user_id = ? OR t.to_user_id = ?)", userID, userID).
		OrderBy(TxDate + " DESC").
		Limit(10)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", mark, zap.Error(err))
		return nil, err
	}

	q := dbClient.Query{
		Name:     "GetOperations",
		QueryRaw: query,
	}

	var operations []model.Operations

	err = r.db.DB().ScanAllContext(ctx, &operations, q, args...)
	if err != nil {
		logger.Error("failed to execute query", mark, zap.Error(err))
		return nil, err
	}

	return operations, nil
}
