// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Jira-Analyzer/backend-services/internal/repository (interfaces: IIssueRepository,IProjectRepository)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . IIssueRepository,IProjectRepository
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/Jira-Analyzer/backend-services/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockIIssueRepository is a mock of IIssueRepository interface.
type MockIIssueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIIssueRepositoryMockRecorder
}

// MockIIssueRepositoryMockRecorder is the mock recorder for MockIIssueRepository.
type MockIIssueRepositoryMockRecorder struct {
	mock *MockIIssueRepository
}

// NewMockIIssueRepository creates a new mock instance.
func NewMockIIssueRepository(ctrl *gomock.Controller) *MockIIssueRepository {
	mock := &MockIIssueRepository{ctrl: ctrl}
	mock.recorder = &MockIIssueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIssueRepository) EXPECT() *MockIIssueRepositoryMockRecorder {
	return m.recorder
}

// GetIssuesByProject mocks base method.
func (m *MockIIssueRepository) GetIssuesByProject(arg0 context.Context, arg1 int) ([]domain.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesByProject", arg0, arg1)
	ret0, _ := ret[0].([]domain.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesByProject indicates an expected call of GetIssuesByProject.
func (mr *MockIIssueRepositoryMockRecorder) GetIssuesByProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesByProject", reflect.TypeOf((*MockIIssueRepository)(nil).GetIssuesByProject), arg0, arg1)
}

// InsertIssue mocks base method.
func (m *MockIIssueRepository) InsertIssue(arg0 context.Context, arg1 domain.Issue) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertIssue", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertIssue indicates an expected call of InsertIssue.
func (mr *MockIIssueRepositoryMockRecorder) InsertIssue(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertIssue", reflect.TypeOf((*MockIIssueRepository)(nil).InsertIssue), arg0, arg1)
}

// UpdateIssue mocks base method.
func (m *MockIIssueRepository) UpdateIssue(arg0 context.Context, arg1 domain.Issue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIssue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIssue indicates an expected call of UpdateIssue.
func (mr *MockIIssueRepositoryMockRecorder) UpdateIssue(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssue", reflect.TypeOf((*MockIIssueRepository)(nil).UpdateIssue), arg0, arg1)
}

// MockIProjectRepository is a mock of IProjectRepository interface.
type MockIProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProjectRepositoryMockRecorder
}

// MockIProjectRepositoryMockRecorder is the mock recorder for MockIProjectRepository.
type MockIProjectRepositoryMockRecorder struct {
	mock *MockIProjectRepository
}

// NewMockIProjectRepository creates a new mock instance.
func NewMockIProjectRepository(ctrl *gomock.Controller) *MockIProjectRepository {
	mock := &MockIProjectRepository{ctrl: ctrl}
	mock.recorder = &MockIProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProjectRepository) EXPECT() *MockIProjectRepositoryMockRecorder {
	return m.recorder
}

// GetProjects mocks base method.
func (m *MockIProjectRepository) GetProjects(arg0 context.Context) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockIProjectRepositoryMockRecorder) GetProjects(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockIProjectRepository)(nil).GetProjects), arg0)
}

// GetProjectsByRange mocks base method.
func (m *MockIProjectRepository) GetProjectsByRange(arg0 context.Context, arg1, arg2 int) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsByRange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsByRange indicates an expected call of GetProjectsByRange.
func (mr *MockIProjectRepositoryMockRecorder) GetProjectsByRange(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsByRange", reflect.TypeOf((*MockIProjectRepository)(nil).GetProjectsByRange), arg0, arg1, arg2)
}

// InsertProject mocks base method.
func (m *MockIProjectRepository) InsertProject(arg0 context.Context, arg1 domain.Project) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProject", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProject indicates an expected call of InsertProject.
func (mr *MockIProjectRepositoryMockRecorder) InsertProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProject", reflect.TypeOf((*MockIProjectRepository)(nil).InsertProject), arg0, arg1)
}

// UpdateProject mocks base method.
func (m *MockIProjectRepository) UpdateProject(arg0 context.Context, arg1 domain.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockIProjectRepositoryMockRecorder) UpdateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockIProjectRepository)(nil).UpdateProject), arg0, arg1)
}
