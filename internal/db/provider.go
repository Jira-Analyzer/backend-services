package provider

import (
	"fmt"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	*sqlx.DB
}

func NewPsqlProvider(config *DbConfig) (*Provider, error) {
	connectionFmt := "postgresql://@%s/%s?user=%s&password=%s&sslmode=disable"
	db, err := sqlx.Open("pgx", fmt.Sprintf(connectionFmt, config.Host, config.Name, config.User, config.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to add database to pool. Error: %w", errorlib.ErrHttpInternal)
	}
	if db.Ping() != nil {
		return nil, fmt.Errorf("failed to ping database. Error: %w", errorlib.ErrHttpInternal)
	}

	return &Provider{
		DB: db,
	}, nil
}
