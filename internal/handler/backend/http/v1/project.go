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
	router.HandleFunc("/projects", handler.getProjectsPage).Methods(http.MethodGet).Queries("limit", "{limit}", "page", "{page}")
	router.HandleFunc("/projects", handler.getAll).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id:[0-9]+}", handler.getProjectById).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id:[0-9]+}/fetch", handler.fetchProject).Methods(http.MethodPatch)
	router.HandleFunc("/projects/fetch", handler.fetchProjects).Methods(http.MethodPatch)
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

func (handler *ProjectHandler) fetchProject(writer http.ResponseWriter, request *http.Request) {
	handler.proxy.ServeHTTP(writer, request)
}

func (handler *ProjectHandler) fetchProjects(writer http.ResponseWriter, request *http.Request) {
	handler.proxy.ServeHTTP(writer, request)
}
