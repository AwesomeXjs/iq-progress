package pg

import (
	"context"

	"github.com/AwesomeXjs/iq-progress/pkg/dbclient"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Client - db client struct
type pgClient struct {
	masterDBC dbclient.DB
}

// New - create new db client with pgxpool.Connect
func New(ctx context.Context, dsn string) (dbclient.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

// DB - returning db
func (c *pgClient) DB() dbclient.DB {
	return c.masterDBC
}

// Close - close db
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
