package v1

import (
	"encoding/json"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	service "github.com/Jira-Analyzer/backend-services/internal/service/connector"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type IssueHandler struct {
	issueService service.IIssueService
}

func NewIssueHandler(issueService service.IIssueService) *IssueHandler {
	return &IssueHandler{
		issueService: issueService,
	}
}

func (handler *IssueHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/issues/create", handler.InsertIssue).Methods(http.MethodPost)
	router.HandleFunc("/issues/update", handler.UpdateIssue).Methods(http.MethodPatch)
}

func (handler *IssueHandler) InsertIssue(w http.ResponseWriter, r *http.Request) {
	var issue domain.Issue

	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedID, err := handler.issueService.InsertIssue(r.Context(), issue)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Failed to create issue", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 4886ddf (12-adding-err-checks)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"id": insertedID})
	if err != nil {
=======
	if err = json.NewEncoder(w).Encode(map[string]interface{}{"id": insertedID}); err != nil {
>>>>>>> 1c2ee09 (12-last-fixes)
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
<<<<<<< HEAD
=======
	json.NewEncoder(w).Encode(map[string]interface{}{"id": insertedID})
>>>>>>> a7c5cc4 (12-removing-getters-adding-inserters)
=======
>>>>>>> 4886ddf (12-adding-err-checks)
}

func (handler *IssueHandler) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issue domain.Issue

	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.issueService.UpdateIssue(r.Context(), issue)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Failed to update issue", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(map[string]interface{}{"message": "Issue updated successfully"}); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Issue updated successfully"})
}
