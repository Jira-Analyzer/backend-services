//go:build unit
// +build unit

package psql

import (
	"context"
	"testing"
	"time"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var issueRows = sqlxmock.NewRows([]string{"id", "project_id", "author_id", "reporter_id", "key", "summary", "type", "priority", "status", "created_time", "closed_time", "updated_time", "time_spent"}).
	AddRow(12330737, 10730, 1, 1, "AGILA-44", "Option to use JSP Includ…g forms for Task nodes.", "Bug", "Major", "Open", time.Time{}, time.Time{}, time.Time{}, 0)

func TestIssueRepository_GetIssuesByProject(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() {
		db.Close()
	})

	mock.ExpectQuery(`^SELECT (.+) FROM "Issue" WHERE project_id=(.+)$`).
		WithArgs(10730).
		WillReturnRows(issueRows)

	repo := NewIssueRepository(&provider.Provider{
		DB: db,
	})

	issues, err := repo.GetIssuesByProject(context.Background(), 10730)
	if assert.NoError(t, err) {
		assert.Len(t, issues, 1)
	}

	issues, err = repo.GetIssuesByProject(context.Background(), 123)
	if assert.Error(t, err) {
		assert.Nil(t, issues)
	}
}
