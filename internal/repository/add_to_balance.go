package repository

import (
	"context"
	"strings"

	"github.com/AwesomeXjs/iq-progress/internal/utils"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) AddToBalance(ctx context.Context, userID int, amount int) (int, error) {

	const mark = "Repository.AddToBalance"

	updateBuilder := sq.Update(UserTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{IDColumn: userID}).
		Set(BalanceColumn, sq.Expr("balance + ?", amount)).
		Suffix(ReturnBalanceColumn)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", mark, zap.Error(err))
		return 0, err
	}

	q := dbClient.Query{
		Name:     "AddToBalance",
		QueryRaw: query,
	}

	var balance int
	err = r.db.DB().ScanOneContext(ctx, &balance, q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return 0, utils.ErrUserNotFound
		}
		logger.Error("failed to execute query", mark, zap.Error(err))
		return 0, err
	}

	return balance, nil
}
