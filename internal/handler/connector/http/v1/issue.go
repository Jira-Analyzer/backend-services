package v1

import (
	"net/http"
	"strconv"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/Jira-Analyzer/backend-services/internal/util"
	"github.com/gorilla/mux"
)

type IssueHandler struct {
	service *service.Service
}

func NewIssueHandler(service *service.Service) *IssueHandler {
	return &IssueHandler{
		service: service,
	}
}

func (handler *IssueHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/issues/fetch", handler.FetchIssueByID).Methods(http.MethodPatch).Queries("project_id", "{project_id}")
}

func (handler *IssueHandler) FetchIssueByID(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(mux.Vars(r)["project_id"])
	if err != nil || projectId < 0 {
		jsonErr := errorlib.GetJSONError("invalid query 'project_id' parameter", errorlib.ErrHttpInvalidRequestData)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	err = handler.service.FetchIssue(projectId)
	if err != nil {
		jsonErr := errorlib.GetJSONError("failed to fetch issue", errorlib.ErrHttpInternal)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	response := map[string]interface{}{"message": "Project updated successfully"}
	util.WriteJSON(w, response)
}
