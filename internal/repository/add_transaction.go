package repository

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/dbClient"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) AddTransaction(ctx context.Context, data *model.TxData, txType string) error {
	insertBuilder := sq.Insert(TransactionTable).
		PlaceholderFormat(sq.Dollar).
		Columns(FromUserIDColumn, ToUserIDColumn, AmountColumn, TypeColumn).
		Values(data.Sender, data.Receiver, data.Amount, txType)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return err
	}

	q := dbClient.Query{
		Name:     "AddTransaction",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
