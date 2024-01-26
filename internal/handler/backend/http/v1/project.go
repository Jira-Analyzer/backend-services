package v1

import (
	"net/http"
	"strconv"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	service "github.com/Jira-Analyzer/backend-services/internal/service/backend"
	"github.com/Jira-Analyzer/backend-services/internal/util"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ProjectHandler struct {
	projectService service.IProjectService
	proxy          *service.ProxyService
}

func NewProjectHandler(service *service.Services) *ProjectHandler {
	return &ProjectHandler{
		projectService: service.ProjectService,
		proxy:          service.Proxy,
	}
}

func (handler *ProjectHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/projects", handler.getProjectsPage).Methods(http.MethodGet, http.MethodOptions).Queries("limit", "{limit}", "page", "{page}")
	router.HandleFunc("/projects", handler.getAll).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/projects/{id:[0-9]+}", handler.getProjectById).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/projects/{id:[0-9]+}/fetch", handler.fetchProject).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/projects/fetch", handler.fetchProjects).Methods(http.MethodPatch, http.MethodOptions).Queries("limit", "{limit}", "page", "{page}")
}

type ProjectsDTO struct {
	Links    []string         `json:"_links"`
	Projects []domain.Project `json:"projects"`
	Count    int              `json:"count"`
	Page     int              `json:"page,omitempty"`
}

type ProjectDTO struct {
	Links   []string       `json:"_links"`
	Project domain.Project `json:"project"`
}

// getAll gets all fetched projects
// @Summary      Get all fetched projects
// @Description  get saved projects
// @Tags         project
// @Accept       json
// @Produce      json
// @Success      200  {opject}  ProjectsDTO
// @Failure      400  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Router       /projects [get]
func (handler *ProjectHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	projects, err := handler.projectService.GetProjects(request.Context())
	if err != nil {
		log.Error(err)

		jsonerr := errorlib.GetJSONError("server failed to fetch projects", err)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	response := ProjectsDTO{
		Projects: projects,
		Count:    len(projects),
	}
	util.WriteJSON(writer, &response)
}

// getProjectsPage gets all projects
// @Summary      Get projects by pages
// @Description  support pagination for ptojects
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        limit   query      int  true  "Max number of projects"
// @Param        page   query      int  true  "Page number"
// @Success      200  {opject}  ProjectsDTO
// @Failure      400  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Router       /projects [get]
func (handler *ProjectHandler) getProjectsPage(writer http.ResponseWriter, request *http.Request) {
	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		jsonerr := errorlib.GetJSONError("invalid query 'limit' parameter", errorlib.ErrHttpInvalidRequestData)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		jsonerr := errorlib.GetJSONError("invalid query 'page' parameter", errorlib.ErrHttpInvalidRequestData)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	projects, err := handler.projectService.GetProjectsByRange(request.Context(), (page-1)*limit, limit)
	if err != nil {
		log.Error(err)

		jsonerr := errorlib.GetJSONError("server failed to fetch projects", err)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}
	response := ProjectsDTO{
		Projects: projects,
		Count:    len(projects),
		Page:     page,
	}
	util.WriteJSON(writer, &response)
}

// getProjectById gets project by id
// @Summary      Get project info by ID
// @Description  get one fetched projects
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Project ID"
// @Success      200  {opject}  ProjectDTO
// @Failure      400  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Router       /projects/{id} [get]
func (handler *ProjectHandler) getProjectById(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	project, err := handler.projectService.GetProjectById(request.Context(), id)
	if err != nil {
		log.Error(err)

		jsonerr := errorlib.GetJSONError("server failed to get project", err)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	response := ProjectDTO{
		Project: *project,
	}
	util.WriteJSON(writer, &response)
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
func (handler *ProjectHandler) fetchProject(writer http.ResponseWriter, request *http.Request) {
	handler.proxy.ServeHTTP(writer, request)
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
func (handler *ProjectHandler) fetchProjects(writer http.ResponseWriter, request *http.Request) {
	handler.proxy.ServeHTTP(writer, request)
}
