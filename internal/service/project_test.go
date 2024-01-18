//go:build unit
// +build unit

package service

import (
	"context"
	"testing"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	mock_repository "github.com/Jira-Analyzer/backend-services/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var projectsList = []domain.Project{
	{},
	{},
	{},
	{},
}

func TestProjectService_GetProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIProjectRepository(ctrl)

	mockRepo.EXPECT().GetProjects(gomock.Any()).AnyTimes().Return(projectsList, nil)

	service := NewProjectService(mockRepo)

	projects, err := service.GetProjects(context.Background())
	if assert.NoError(t, err) {
		assert.Len(t, projects, 4)
		assert.Equal(t, projectsList, projects)
	}
}

func TestProjectService_GetProjectsByRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIProjectRepository(ctrl)

	mockRepo.EXPECT().GetProjectsByRange(gomock.Any(), 0, 4).AnyTimes().Return(projectsList, nil)

	service := NewProjectService(mockRepo)

	projects, err := service.GetProjectsByRange(context.Background(), 0, 4)
	if assert.NoError(t, err) {
		assert.Len(t, projects, 4)
		assert.Equal(t, projectsList, projects)
	}
}
