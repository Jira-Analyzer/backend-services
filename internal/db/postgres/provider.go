package provider

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	*sqlx.DB
}

func NewProvider(config *Config) (*Provider, error) {
	connectionFmt := "postgresql://@%s/%s?user=%s&password=%s&sslmode=disable"
	db, err := sqlx.Open("pgx", fmt.Sprintf(connectionFmt, config.Host, config.Name, config.User, config.Password))
	if err != nil {
		return nil, fmt.Errorf("Failed to add database to pool. Error: %w", err)
	}

	return &Provider{
		DB: db,
	}, nil
}

func (provider *Provider) Close() error {
	if provider != nil {
		return provider.DB.Close()
	}
	return nil
}
