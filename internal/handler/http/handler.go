package http

import (
	v1 "github.com/Jira-Analyzer/backend-services/internal/handler/http/v1"
	"github.com/Jira-Analyzer/backend-services/internal/service"
)

type Handler struct {
	v1 *v1.Handler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		v1: v1.NewHandler(services),
	}
}
