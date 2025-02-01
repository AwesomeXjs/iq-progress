package repository

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/internal/model"
	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	sq "github.com/Masterminds/squirrel"
)

// AddTransaction adds a new transaction record to the database with the specified data and transaction type.
func (r *Repository) AddTransaction(ctx context.Context, data *model.TxData, txType string) error {
	insertBuilder := sq.Insert(TransactionTable).
		PlaceholderFormat(sq.Dollar).
		Columns(FromUserIDColumn, ToUserIDColumn, AmountColumn, TypeColumn).
		Values(data.Sender, data.Receiver, data.Amount, txType)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return err
	}

	q := dbclient.Query{
		Name:     "AddTransaction",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
