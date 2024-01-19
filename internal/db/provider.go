package provider

import (
	"fmt"

	"github.com/Jira-Analyzer/backend-services/internal/config"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	*sqlx.DB
}

func NewPsqlProvider(config *config.DbConfig) (*Provider, error) {
	connectionFmt := "postgresql://@%s/%s?user=%s&password=%s&sslmode=disable"
	db, err := sqlx.Open("pgx", fmt.Sprintf(connectionFmt, config.Host, config.Name, config.User, config.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to add database to pool. Error: %w", errorlib.InternalError)
	}
	if db.Ping() != nil {
		return nil, fmt.Errorf("failed to ping database. Error: %w", errorlib.InternalError)
	}

	return &Provider{
		DB: db,
	}, nil
}
