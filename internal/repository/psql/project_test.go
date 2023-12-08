package psql

import (
	"testing"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var projectRows = sqlxmock.NewRows([]string{"id", "name", "description", "avatar_url", "type", "archived"}).
	AddRow(10730, "AGILA", "some description", "https://issues.apache.or…pid=10730&avatarId=10011", "software", false).
	AddRow(123, "prjc2", "some", "https://issues.apache.or…pid=10730&avatarId=10011", "software", true)

func TestIssueRepository_GetProjects(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT (.+) FROM (.+)$`).
		WillReturnRows(projectRows)

	repo := NewProjectRepository(&provider.Provider{
		DB: db,
	})

	issues, err := repo.GetProjects()
	assert.Nil(t, err)
	assert.Len(t, issues, 2)
}
