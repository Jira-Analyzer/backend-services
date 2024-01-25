package v1

import (
	"net/http"
	"strconv"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/Jira-Analyzer/backend-services/internal/util"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ProjectHandler struct {
	service *service.Service
}

func NewProjectHandler(service *service.Service) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (handler *ProjectHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/projects/fetch", handler.FetchAllProjects).Methods(http.MethodPatch).Queries("limit", "{limit}", "page", "{page}")
	router.HandleFunc("/projects/{id:[0-9]+}/fetch", handler.FetchProjectByID).Methods(http.MethodPatch)
}

func (handler *ProjectHandler) FetchAllProjects(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		jsonErr := errorlib.GetJSONError("invalid query 'limit' parameter", errorlib.ErrHttpInvalidRequestData)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		jsonErr := errorlib.GetJSONError("invalid query 'page' parameter", errorlib.ErrHttpInvalidRequestData)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	projects, err := handler.service.FetchProjects(page, limit)
	if err != nil {
		logrus.Error(err)
		jsonErr := errorlib.GetJSONError("failed to fetch projects", errorlib.ErrHttpInternal)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	util.WriteJSON(w, projects)
}

func (handler *ProjectHandler) FetchProjectByID(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := handler.service.FetchProject(projectId)
	if err != nil {
		logrus.Error(err)
		jsonErr := errorlib.GetJSONError("failed to fetch project", errorlib.ErrHttpInternal)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	response := map[string]interface{}{"message": "Project updated successfully"}
	util.WriteJSON(w, response)
}
