package v1

import (
	"net/http"
	"strconv"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	service "github.com/Jira-Analyzer/backend-services/internal/service/backend"
	"github.com/Jira-Analyzer/backend-services/internal/util"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type IssueHandler struct {
	service service.IIssueService
	proxy   *service.ProxyService
}

func NewIssueHandler(service *service.Services) *IssueHandler {
	return &IssueHandler{
		service: service.IssueService,
		proxy:   service.Proxy,
	}
}

func (handler *IssueHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/issues", handler.getIssuesByProjectId).Methods(http.MethodGet, http.MethodOptions).Queries("project_id", "{project_id}")
	router.HandleFunc("/issues/statistics", handler.getStatistics).Methods(http.MethodGet, http.MethodOptions).Queries("project_id", "{project_id}")
	router.HandleFunc("/issues/fetch", handler.fetchIssues).Methods(http.MethodPatch, http.MethodOptions).Queries("project_id", "{project_id}")
}

type issuesDTO struct {
	Links     []string       `json:"_links"`
	ProjectId int            `json:"project_id"`
	Issues    []domain.Issue `json:"issues"`
}

// getIssuesByProjectId gets issues by project id
// @Summary      Get issues list by project id
// @Description  get list by ID
// @Tags         issue
// @Accept       json
// @Produce      json
// @Param        project_id   query      int  true  "Project ID"
// @Success      200  {object}  issuesDTO
// @Failure      400  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Router       /issues [get]
func (handler *IssueHandler) getIssuesByProjectId(writer http.ResponseWriter, request *http.Request) {
	projectId, err := strconv.Atoi(mux.Vars(request)["project_id"])
	if err != nil || projectId < 0 {
		jsonerr := errorlib.GetJSONError("invalid query 'project_id' parameter", errorlib.ErrHttpInvalidRequestData)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	issues, err := handler.service.GetIssuesByProject(request.Context(), projectId)
	if err != nil {
		logrus.Error(err)

		jsonerr := errorlib.GetJSONError("server failed to fetch issues", err)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	response := issuesDTO{
		ProjectId: projectId,
		Issues:    issues,
	}
	util.WriteJSON(writer, &response)
}

type statisticsDTO struct {
	Links      []string `json:"_links"`
	ProjectId  int      `json:"project_id"`
	Statistics struct {
		Total         int `json:"total_number"`
		Opened        int `json:"opened_number"`
		Closed        int `json:"closed_number"`
		AverageTime   int `json:"average_time"`
		OpenedAverage int `json:"weekly_average_opened"`
	} `json:"statistics"`
}

// getStatistics gets issues by project id
// @Summary      Get issues statistics on project
// @Description  get statistics by project ID
// @Tags         issue
// @Accept       json
// @Produce      json
// @Param        project_id   query      int  true  "Project ID"
// @Success      200  {object}  statisticsDTO
// @Failure      400  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Router       /issues/statistics [get]
func (handler *IssueHandler) getStatistics(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil || id < 0 {
		jsonerr := errorlib.GetJSONError("invalid query 'project_id' parameter", errorlib.ErrHttpInvalidRequestData)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}
	issues, err := handler.service.GetIssuesByProject(request.Context(), id)
	if err != nil {
		log.Error(err)

		jsonerr := errorlib.GetJSONError("server failed to get issues", err)
		writer.WriteHeader(jsonerr.Error.Code)
		util.WriteJSON(writer, &jsonerr)
		return
	}

	statistics := statisticsDTO{}
	statistics.Statistics.Total = len(issues)
	statistics.Statistics.Opened = len(handler.service.FilterOpenedIssues(request.Context(), issues))
	statistics.Statistics.Closed = len(handler.service.FilterClosedIssues(request.Context(), issues))
	statistics.Statistics.AverageTime = handler.service.GetWeekAverageCreatedNumber(request.Context(), issues)
	statistics.Statistics.OpenedAverage = handler.service.GetWeekAverageCreatedNumber(request.Context(), issues)

	util.WriteJSON(writer, &statistics)
}

// fetchIssues save issues from Jira to db
// @Summary      Fetch project's issues locally
// @Description  fetch issues
// @Tags         issue
// @Accept       json
// @Produce      json
// @Param        project_id   query      int  true  "Project ID"
// @Success      200  {string}  Success
// @Failure      400  {object}  errorlib.JSONError
// @Failure      408  {object}  errorlib.JSONError
// @Failure      409  {object}  errorlib.JSONError
// @Failure      422  {object}  errorlib.JSONError
// @Failure      500  {object}  errorlib.JSONError
// @Failure      504  {object}  errorlib.JSONError
// @Router       /issues/fetch [patch]
func (handler *IssueHandler) fetchIssues(writer http.ResponseWriter, request *http.Request) {
	handler.proxy.ServeHTTP(writer, request)
}
