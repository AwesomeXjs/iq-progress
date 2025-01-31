package config

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/iq-progress/pkg/logger"
	"go.uber.org/zap"
)

const (
	// PgDsn is the environment variable key for the PostgreSQL Data Source Name (DSN).
	// It should be used to fetch the DSN from the environment, typically specified in the .env file.
	PgDsn = "PG_DSN"
)

// PgConfig implements the PGConfig interface, storing the DSN.
type PgConfig struct {
	dsn string
}

// NewPgConfig creates a new PGConfig instance by reading the DSN from environment variables.
// It returns an error if the DSN is not set.
func NewPgConfig() (*PgConfig, error) {
	const mark = "Config.PGConfig"
	dsn := os.Getenv(PgDsn)
	if len(dsn) == 0 {
		logger.Error("failed to get db dsn", mark, zap.String("db dsn", PgDsn))
		return nil, fmt.Errorf("env %s is empty", PgDsn)
	}

	return &PgConfig{
		dsn: dsn,
	}, nil
}

// GetDSN returns the database connection string (DSN) from the pgConfig instance.
func (p *PgConfig) GetDSN() string {
	return p.dsn
}
