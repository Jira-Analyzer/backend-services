package http

import (
	"github.com/Jira-Analyzer/backend-services/internal/handler/connector/http/v1"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/gorilla/mux"
)

type Handler struct {
	v1 *v1.Handler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		v1: v1.NewHandler(services),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/api").Subrouter()
	handler.v1.SetRouter(sub)
}
