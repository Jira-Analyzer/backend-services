package v1

import (
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/gorilla/mux"
)

type Handler struct {
	projectHandler *ProjectHandler
	issueHandler   *IssueHandler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		projectHandler: NewProjectHandler(services.ProjectService),
		issueHandler:   NewIssueHandler(services.IssueService),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/v1").Subrouter()
	handler.projectHandler.SetRouter(sub)
}
