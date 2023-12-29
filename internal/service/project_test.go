package service

import (
	"testing"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/repository/psql"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestProjectService_GetProjects(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	projectRows := sqlxmock.NewRows([]string{"id", "name", "description", "avatar_url", "type", "archived"}).
		AddRow(10730, "AGILA", "some description", "https://issues.apache.or…pid=10730&avatarId=10011", "software", false).
		AddRow(123, "prjc2", "some", "https://issues.apache.or…pid=10730&avatarId=10011", "software", true)

	mock.ExpectQuery(`^SELECT (.+) FROM "Project"$`).
		WillReturnRows(projectRows)

	repo := psql.NewProjectRepository(&provider.Provider{
		DB: db,
	})

	service := NewProjectService(repo)

	projects, err := service.GetProjects()
	if assert.NoError(t, err) {
		assert.Len(t, projects, 2)
	}
}

func TestProjectService_GetNFirstProjects(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	projectRows := sqlxmock.NewRows([]string{"id", "name", "description", "avatar_url", "type", "archived"}).
		AddRow(10730, "AGILA", "some description", "https://issues.apache.or…pid=10730&avatarId=10011", "software", false).
		AddRow(123, "prjc2", "some", "https://issues.apache.or…pid=10730&avatarId=10011", "software", true)

	mock.ExpectQuery(`^SELECT (.+) FROM "Project"$`).
		WillReturnRows(projectRows)

	repo := psql.NewProjectRepository(&provider.Provider{
		DB: db,
	})

	service := NewProjectService(repo)

	projects, err := service.GetNFirstProjects(1)
	if assert.NoError(t, err) {
		assert.Len(t, projects, 1)
		assert.Equal(t, projects[0].Name, "AGILA")
	}

	projects, err = service.GetNFirstProjects(-1)
	if assert.Error(t, err) {
		assert.Nil(t, projects)
	}

	projects, err = service.GetNFirstProjects(100)
	if assert.Error(t, err) {
		assert.Nil(t, projects)
	}
}
