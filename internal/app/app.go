package app

import (
	"fmt"

	"github.com/Jira-Analyzer/backend-services/internal/config"
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	handler "github.com/Jira-Analyzer/backend-services/internal/handler/http"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
	"github.com/Jira-Analyzer/backend-services/internal/server"
	"github.com/Jira-Analyzer/backend-services/internal/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type App struct {
	server   *server.Server
	provider *provider.Provider
}

func NewApp(config *config.Config) *App {
	provider, err := provider.NewPsqlProvider(&config.DbConfig)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to initialize db with error: %w", err))
	}

	repos := repository.NewRepositories(provider)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	router := mux.NewRouter()
	handlers.SetRouter(router)

	server := server.NewServer(&config.ServerConfig, router)

	return &App{
		server:   server,
		provider: provider,
	}
}

func (app *App) Start() {
	app.server.Start()
}

func (app *App) GetServerNotify() <-chan error {
	return app.server.GetNotify()
}

func (app *App) Stop() {
	err := app.server.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("Graceful shutdown was ended with error: %w", err))
	}

	err = app.provider.Close()
	if err != nil {
		log.Error(fmt.Errorf("Failed to close db connection with error: %w", err))
	}
}
