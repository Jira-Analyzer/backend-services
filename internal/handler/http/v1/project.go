package v1

import (
	"net/http"

	"github.com/Jira-Analyzer/backend-services/internal/service"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (handler *ProjectHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/projects", handler.getAll).Methods(http.MethodGet)
}

func (handler *ProjectHandler) getAll(writer http.ResponseWriter, request *http.Request) {

}
