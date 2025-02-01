package pg

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key string

// TxKey is the context key for database transactions.
const (
	TxKey key = "tx"
)

// pg represents a PostgreSQL database client using pgxpool.
type pg struct {
	dbc *pgxpool.Pool // Database connection pool
}

// NewDB initializes and returns a new database client.
func NewDB(dbc *pgxpool.Pool) dbclient.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOneContext executes a query and scans the first result into the destination.
func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q dbclient.Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAllContext executes a query and scans all results into the destination.
func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q dbclient.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// ExecContext executes a query without returning any rows.
func (p *pg) ExecContext(ctx context.Context, q dbclient.Query, args ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext executes a query and returns multiple rows.
func (p *pg) QueryContext(ctx context.Context, q dbclient.Query, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext executes a query and returns a single row.
func (p *pg) QueryRowContext(ctx context.Context, q dbclient.Query, args ...interface{}) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx starts a new database transaction with the given options.
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping checks the connection to the database.
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close closes the database connection pool.
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx adds a transaction to the context.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
