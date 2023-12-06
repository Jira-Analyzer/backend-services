package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ProjectHandler struct {
	projectService service.IProjectService
}

func NewProjectHandler(service service.IProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

func (handler *ProjectHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/projects/create", handler.InsertProject).Methods(http.MethodPost)
	router.HandleFunc("/projects/update", handler.UpdateProject).Methods(http.MethodPatch)
}

func (handler *ProjectHandler) InsertProject(w http.ResponseWriter, r *http.Request) {
	var project domain.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedID, err := handler.projectService.InsertProject(r.Context(), project)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(map[string]interface{}{"id": insertedID}); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"id": insertedID})
}

func (handler *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var project domain.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.projectService.UpdateProject(r.Context(), project)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(map[string]interface{}{"message": "Project updated successfully"}); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Project updated successfully"})
}
