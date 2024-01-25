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
}

func NewIssueHandler(service service.IIssueService) *IssueHandler {
	return &IssueHandler{
		service: service,
	}
}

func (handler *IssueHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/issues", handler.getIssuesByProjectId).Methods(http.MethodGet).Queries("project_id", "{project_id}")
	router.HandleFunc("/issues/statistics", handler.getStatistics).Methods(http.MethodGet).Queries("project_id", "{project_id}")
}

type issuesDTO struct {
	Links     []string       `json:"_links"`
	ProjectId int            `json:"project_id"`
	Issues    []domain.Issue `json:"issues"`
}

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
