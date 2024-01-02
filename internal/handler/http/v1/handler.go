package v1

import "github.com/Jira-Analyzer/backend-services/internal/service"

type Handler struct {
	projectHandler *ProjectHandler
}

func NewHandler(serivces *service.Services) *Handler {
	return &Handler{
		projectHandler: NewProjectHandler(serivces.ProjectService),
	}
}
