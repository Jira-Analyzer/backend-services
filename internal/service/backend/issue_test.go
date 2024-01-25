//go:build unit
// +build unit

package service

import (
	"context"
	"testing"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	mock_repository "github.com/Jira-Analyzer/backend-services/internal/repository/mock"
	"github.com/Jira-Analyzer/backend-services/internal/repository/psql"
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

func TestIssueService_FilterOpenedIssues(t *testing.T) {
	service := NewIssueService(&psql.IssueRepository{})
	data := []domain.Issue{
		{Status: "Opened"},
		{Status: "Opened"},
		{Status: "Opened"},
		{Status: "Closed"},
		{Status: "Closed"},
	}
	empty := []domain.Issue{}

	filtered := service.FilterOpenedIssues(context.Background(), data)
	assert.Len(t, filtered, 3)
	filtered = service.FilterOpenedIssues(context.Background(), empty)
	assert.Len(t, filtered, 0)
}

func TestIssueService_FilterClosedIssues(t *testing.T) {
	service := NewIssueService(&psql.IssueRepository{})
	data := []domain.Issue{
		{Status: "Opened"},
		{Status: "Opened"},
		{Status: "Opened"},
		{Status: "Closed"},
		{Status: "Closed"},
	}
	empty := []domain.Issue{}

	filtered := service.FilterClosedIssues(context.Background(), data)
	assert.Len(t, filtered, 2)
	filtered = service.FilterClosedIssues(context.Background(), empty)
	assert.Len(t, filtered, 0)
}

func TestIssueService_GetAverageTimeSpent(t *testing.T) {
	service := NewIssueService(&psql.IssueRepository{})
	data := []domain.Issue{
		{TimeSpent: 5 * time.Hour},
		{TimeSpent: 10 * time.Hour},
		{TimeSpent: 5 * time.Hour},
		{TimeSpent: 10 * time.Hour},
		{TimeSpent: 5 * time.Hour},
	}
	empty := []domain.Issue{}

	duration := service.GetAverageTimeSpent(context.Background(), data)
	assert.Equal(t, duration, 7*time.Hour)
	duration = service.GetAverageTimeSpent(context.Background(), empty)
	assert.Equal(t, duration, time.Duration(0))
}

func TestIssueService_GetWeekAverageCreatedNumber(t *testing.T) {
	service := NewIssueService(&psql.IssueRepository{})
	now := time.Now()
	data := []domain.Issue{
		{CreatedTime: now.AddDate(0, 2, -7)},
		{CreatedTime: now.AddDate(-2, -1, -7)},
		{CreatedTime: now.AddDate(0, 0, -1)},
		{CreatedTime: now.AddDate(0, 0, -2)},
		{CreatedTime: now.AddDate(0, 0, -1)},
		{CreatedTime: now.AddDate(0, 0, -5)},
		{CreatedTime: now.AddDate(0, 0, -1)},
		{CreatedTime: now.AddDate(0, 0, -2)},
		{CreatedTime: now.AddDate(0, 0, -1)},
		{CreatedTime: now.AddDate(0, 0, -5)},
		{CreatedTime: now.AddDate(2, -1, -7)},
		{CreatedTime: now.AddDate(0, 0, 1)},
	}
	empty := []domain.Issue{}

	num := service.GetWeekAverageCreatedNumber(context.Background(), data)
	assert.Equal(t, 8/7, num)
	num = service.GetWeekAverageCreatedNumber(context.Background(), empty)
	assert.Equal(t, 0, num)
}
