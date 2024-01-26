package connector

import (
	"fmt"
	"net/http"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	handler "github.com/Jira-Analyzer/backend-services/internal/handler/connector/http"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
	server "github.com/Jira-Analyzer/backend-services/internal/server/backend"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	server   *server.Server
	provider *provider.Provider
}

func NewApp(config *Config, notify chan error) (*App, error) {
	provider, err := provider.NewPsqlProvider(config.DbConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize db with error: %w", err)
	}

	repos := repository.NewRepositories(provider)
	services := service.NewService(config.ConnectorConfig, repos)
	handlers := handler.NewHandler(services)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Set-Cookie", "User-Agent", "Origin"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	router := mux.NewRouter()
	handlers.SetRouter(router)

	server := server.NewServer(&server.ServerConfig{
		Host:         &config.ConnectorConfig.Host,
		ReadTimeout:  &config.ConnectorConfig.ReadTimeout,
		WriteTimeout: &config.ConnectorConfig.WriteTimeout,
	}, c.Handler(router), notify)

	return &App{
		server:   server,
		provider: provider,
	}, nil
}

func (app *App) Start() {
	app.server.Start()
}

func (app *App) Stop() error {
	serverErr := app.server.Stop()
	providerErr := app.provider.Close()
	if serverErr != nil || providerErr != nil {
		return fmt.Errorf("Provider error: %w. Server error: %w", providerErr, serverErr)
	}
	return nil
}
