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

type IssueHandler struct {
	service *service.Service
}

func NewIssueHandler(service *service.Service) *IssueHandler {
	return &IssueHandler{
		service: service,
	}
}

func (handler *IssueHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/issues/fetch", handler.FetchIssueByID).Methods(http.MethodPatch, http.MethodOptions).Queries("project_id", "{project_id}")
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
		logrus.Error(err)
		jsonErr := errorlib.GetJSONError("failed to fetch issue", err)
		w.WriteHeader(jsonErr.Error.Code)
		util.WriteJSON(w, &jsonErr)
		return
	}

	response := map[string]interface{}{"message": "Issues updated successfully"}
	util.WriteJSON(w, response)
}
