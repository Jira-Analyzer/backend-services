package v1

import (
	"github.com/Jira-Analyzer/backend-services/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	projectHandler *ProjectHandler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		projectHandler: NewProjectHandler(services.ProjectService),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/v1").Subrouter()
	handler.projectHandler.SetRouter(sub)
}
