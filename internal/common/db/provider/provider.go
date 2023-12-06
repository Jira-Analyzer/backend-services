package provider

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Provider struct {
	DB *sqlx.DB
}

func NewProvider(config *Config) (*Provider, error) {
	connectionFmt := "host=%s user=%s dbname=%s sslmode=disable password=%s"
	db, err := sqlx.Open("postgres", fmt.Sprintf(connectionFmt, config.Host, config.User, config.Name, config.Password))
	if err != nil {
		return nil, fmt.Errorf("Failed to add database to pool. Error: %s", err.Error())
	}

	return &Provider{
		DB: db,
	}, nil
}

func (provider *Provider) Close() {
	provider.DB.Close()
}
