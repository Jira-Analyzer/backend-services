package jira

import "github.com/Jira-Analyzer/backend-services/internal/domain"

type FetchProjectResponseDTO struct {
	Projects []*domain.Project `json:"projects"`
}

type FetchIssueResponseDTO struct {
	Issues []*domain.Issue `json:"issues"`
}
