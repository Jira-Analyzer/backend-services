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
	router.HandleFunc("/projects/fetch", handler.FetchAllProjects).Methods(http.MethodPatch, http.MethodOptions).Queries("limit", "{limit}", "page", "{page}")
	router.HandleFunc("/projects/{id:[0-9]+}/fetch", handler.FetchProjectByID).Methods(http.MethodPatch, http.MethodOptions)
}

// fetchProjects gets list of all projects from jira
// @Summary      Get short project info by pages
// @Description  support pagination for ptojects
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        limit   query      int  true  "Max number of projects"
// @Param        page   query      int  true  "Page number"
// @Success      200  {object}  dto.ProjectsResponse
// @Failure      400  {object}  errorlib.JSONError
// @Failure      408  {object}  errorlib.JSONError
// @Failure      409  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Failure      504  {object}  errorlib.JSONError
// @Router       /projects/fetch [patch]
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
		jsonErr := errorlib.GetJSONError("failed to fetch projects", err)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	util.WriteJSON(w, projects)
}

// fetchProject save project from Jira to db
// @Summary      Fetch project locally
// @Description  fetch project
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Project ID"
// @Success      200  {string}  Success
// @Failure      400  {object}  errorlib.JSONError
// @Failure      408  {object}  errorlib.JSONError
// @Failure      409  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Failure      504  {object}  errorlib.JSONError
// @Router       /projects/{id}/fetch [patch]
func (handler *ProjectHandler) FetchProjectByID(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := handler.service.FetchProject(projectId)
	if err != nil {
		logrus.Error(err)
		jsonErr := errorlib.GetJSONError("failed to fetch project", err)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	response := map[string]interface{}{"message": "Project updated successfully"}
	util.WriteJSON(w, response)
}
