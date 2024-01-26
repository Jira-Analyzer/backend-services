package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Jira-Analyzer/backend-services/internal/client/jira/dto"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
)

func (s *Service) FetchProjects(page, count int) (*dto.ProjectsResponse, error) {
	projects, err := s.client.FetchProjects(page, count)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch projects from Jira: %w", errorlib.ErrHttpInternal)
	}
	return projects, nil
}

func (s *Service) FetchProject(id int) error {
	project, err := s.client.FetchProject(id)
	if err != nil {
		return fmt.Errorf("failed to fetch project details from Jira: %w", errorlib.ErrHttpInternal)
	}

	dbErr := s.projectsRepo.InsertProject(context.Background(), project)
	if errors.Is(dbErr, sql.ErrNoRows) {
		return fmt.Errorf("project with key '%c' already exists: %w", id, errorlib.ErrHttpConflict)
	} else if err != nil {
		return fmt.Errorf("failed to check project existence in database: %w", errorlib.ErrHttpInternal)
	}
	return nil
}
