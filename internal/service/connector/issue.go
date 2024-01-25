package service

import (
	"context"
	"fmt"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
)

func (s *Service) FetchIssue(projectId int) error {
	issues, err := s.client.FetchIssues(projectId, 100)
	if err != nil {
		return fmt.Errorf("failed to fetch issues from Jira: %w", errorlib.ErrHttpInternal)
	}

	for _, issue := range issues {
		if err := s.issuesRepo.InsertIssue(context.Background(), issue.ToDomainIssue(projectId)); err != nil {
			return fmt.Errorf("failed to insert issue into database: %w", errorlib.ErrHttpConflict)
		}
	}

	return nil
}
