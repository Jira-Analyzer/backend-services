package v1

import (
	service "github.com/Jira-Analyzer/backend-services/internal/service/backend"
	"github.com/gorilla/mux"
)

type Handler struct {
	projectHandler *ProjectHandler
	issueHandler   *IssueHandler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		projectHandler: NewProjectHandler(services),
		issueHandler:   NewIssueHandler(services),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/v1").Subrouter()
	handler.projectHandler.SetRouter(sub)
	handler.issueHandler.SetRouter(sub)
}
