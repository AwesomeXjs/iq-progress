package repository

import (
	"context"
	"strings"

	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// RemoveFromBalance removes a specified amount from a user's balance and returns the updated balance.
func (r *Repository) RemoveFromBalance(ctx context.Context, userID int, amount int) (int, error) {
	const mark = "Repository.RemoveFromBalance"

	updateBuilder := sq.Update(UserTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{IDColumn: userID}).
		Set(BalanceColumn, sq.Expr("balance - ?", amount)).
		Suffix(ReturnBalanceColumn)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", mark, zap.Error(err))
		return 0, err
	}

	q := dbclient.Query{
		Name:     "RemoveFromBalance",
		QueryRaw: query,
	}

	var balance int
	err = r.db.DB().ScanOneContext(ctx, &balance, q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return 0, utils.ErrSenderNotFound
		}
		logger.Error("failed to execute query", mark, zap.Error(err))
		return 0, err
	}

	if balance < 0 {
		return 0, utils.ErrNotEnoughBalance
	}

	return balance, nil
}
