package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jira-Analyzer/backend-services/internal/service"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service service.IProjectService
}

func NewProjectHandler(service service.IProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (handler *ProjectHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/project", handler.getAll).Methods(http.MethodGet)
}

func (handler *ProjectHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	projects, err := handler.service.GetProjects(context.Background())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(&projects)
	if err != nil {
		fmt.Println("uupds")
	}
}
