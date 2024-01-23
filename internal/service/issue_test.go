//go:build unit
// +build unit

package service

import (
	"context"
	"testing"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	mock_repository "github.com/Jira-Analyzer/backend-services/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var issueList = []domain.Issue{
	{ProjectId: 123},
	{ProjectId: 123},
	{ProjectId: 123},
	{ProjectId: 123},
	{ProjectId: 123},
}

func TestIssueRepository_GetIssuesByProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIIssueRepository(ctrl)

	mockRepo.EXPECT().GetIssuesByProject(gomock.Any(), 123).AnyTimes().Return(issueList, nil)
	mockRepo.EXPECT().GetIssuesByProject(gomock.Any(), gomock.Not(123)).AnyTimes().Return(nil, errorlib.ErrHttpInternal)

	service := NewIssueService(mockRepo)

	issues, err := service.GetIssuesByProject(context.Background(), 123)
	if assert.NoError(t, err) {
		assert.Equal(t, issues, issueList)
	}

	issues, err = service.GetIssuesByProject(context.Background(), 112)
	if assert.Error(t, err) {
		assert.Nil(t, issues)
	}
}
