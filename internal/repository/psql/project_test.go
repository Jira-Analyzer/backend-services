//go:build unit
// +build unit

package psql

import (
	"context"
	"testing"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestIssueRepository_GetProjects(t *testing.T) {
	var projectRows = sqlxmock.NewRows([]string{"id", "name", "description", "avatar_url", "type", "archived"}).
		AddRow(1, "AGILA", "some description", "https://issues.apache.or…pid=10730&avatarId=10011", "software", false).
		AddRow(2, "prjc2", "some", "https://issues.apache.or…pid=10730&avatarId=10011", "software", true)

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() {
		db.Close()
	})

	mock.ExpectQuery(`^SELECT (.+) FROM "Project"$`).
		WillReturnRows(projectRows)

	repo := NewProjectRepository(&provider.Provider{
		DB: db,
	})

	projects, err := repo.GetProjects(context.Background())
	if assert.NoError(t, err) {
		assert.Len(t, projects, 2)
	}
}

func TestProjectRepository_GetProjectsByRange(t *testing.T) {
	var projectRows = sqlxmock.NewRows([]string{"id", "name", "description", "avatar_url", "type", "archived"}).
		AddRow(1, "AGILA", "some description", "https://issues.apache.or…pid=10730&avatarId=10011", "software", false).
		AddRow(2, "prjc2", "some", "https://issues.apache.or…pid=10730&avatarId=10011", "software", true)

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() {
		db.Close()
	})

	mock.ExpectQuery(`^SELECT (.+) FROM "Project" ORDER BY id OFFSET (.+) ROWS FETCH FIRST (.+) ROW ONLY$`).
		WithArgs(0, 2).
		WillReturnRows(projectRows)

	repo := NewProjectRepository(&provider.Provider{
		DB: db,
	})

	projects, err := repo.GetProjectsByRange(context.Background(), 0, 2)
	if assert.NoError(t, err) {
		assert.Len(t, projects, 2)
	}
}
