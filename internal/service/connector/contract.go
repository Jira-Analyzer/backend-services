package service

import (
	"github.com/Jira-Analyzer/backend-services/internal/client/jira/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IService

type IService interface {
	FetchProjects(page int, count int) (*dto.ProjectsResponse, error)
	FetchProject(id int) error
	FetchIssue(projectId int) error
}
