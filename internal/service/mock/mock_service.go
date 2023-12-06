// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Jira-Analyzer/backend-services/internal/service (interfaces: IIssueService,IProjectService)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IIssueService,IProjectService
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/Jira-Analyzer/backend-services/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockIIssueService is a mock of IIssueService interface.
type MockIIssueService struct {
	ctrl     *gomock.Controller
	recorder *MockIIssueServiceMockRecorder
}

// MockIIssueServiceMockRecorder is the mock recorder for MockIIssueService.
type MockIIssueServiceMockRecorder struct {
	mock *MockIIssueService
}

// NewMockIIssueService creates a new mock instance.
func NewMockIIssueService(ctrl *gomock.Controller) *MockIIssueService {
	mock := &MockIIssueService{ctrl: ctrl}
	mock.recorder = &MockIIssueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIssueService) EXPECT() *MockIIssueServiceMockRecorder {
	return m.recorder
}

// GetIssuesByProject mocks base method.
func (m *MockIIssueService) GetIssuesByProject(arg0 context.Context, arg1 int64) ([]domain.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesByProject", arg0, arg1)
	ret0, _ := ret[0].([]domain.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesByProject indicates an expected call of GetIssuesByProject.
func (mr *MockIIssueServiceMockRecorder) GetIssuesByProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesByProject", reflect.TypeOf((*MockIIssueService)(nil).GetIssuesByProject), arg0, arg1)
}

// MockIProjectService is a mock of IProjectService interface.
type MockIProjectService struct {
	ctrl     *gomock.Controller
	recorder *MockIProjectServiceMockRecorder
}

// MockIProjectServiceMockRecorder is the mock recorder for MockIProjectService.
type MockIProjectServiceMockRecorder struct {
	mock *MockIProjectService
}

// NewMockIProjectService creates a new mock instance.
func NewMockIProjectService(ctrl *gomock.Controller) *MockIProjectService {
	mock := &MockIProjectService{ctrl: ctrl}
	mock.recorder = &MockIProjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProjectService) EXPECT() *MockIProjectServiceMockRecorder {
	return m.recorder
}

// GetProjects mocks base method.
func (m *MockIProjectService) GetProjects(arg0 context.Context) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockIProjectServiceMockRecorder) GetProjects(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockIProjectService)(nil).GetProjects), arg0)
}

// GetProjectsByRange mocks base method.
func (m *MockIProjectService) GetProjectsByRange(arg0 context.Context, arg1, arg2 int) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsByRange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsByRange indicates an expected call of GetProjectsByRange.
func (mr *MockIProjectServiceMockRecorder) GetProjectsByRange(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsByRange", reflect.TypeOf((*MockIProjectService)(nil).GetProjectsByRange), arg0, arg1, arg2)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Jira-Analyzer/backend-services/internal/service (interfaces: IIssueService,IProjectService)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IIssueService,IProjectService
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/Jira-Analyzer/backend-services/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockIIssueService is a mock of IIssueService interface.
type MockIIssueService struct {
	ctrl     *gomock.Controller
	recorder *MockIIssueServiceMockRecorder
}

// MockIIssueServiceMockRecorder is the mock recorder for MockIIssueService.
type MockIIssueServiceMockRecorder struct {
	mock *MockIIssueService
}

// NewMockIIssueService creates a new mock instance.
func NewMockIIssueService(ctrl *gomock.Controller) *MockIIssueService {
	mock := &MockIIssueService{ctrl: ctrl}
	mock.recorder = &MockIIssueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIssueService) EXPECT() *MockIIssueServiceMockRecorder {
	return m.recorder
}

// GetIssuesByProject mocks base method.
func (m *MockIIssueService) GetIssuesByProject(arg0 context.Context, arg1 int) ([]domain.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesByProject", arg0, arg1)
	ret0, _ := ret[0].([]domain.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesByProject indicates an expected call of GetIssuesByProject.
func (mr *MockIIssueServiceMockRecorder) GetIssuesByProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesByProject", reflect.TypeOf((*MockIIssueService)(nil).GetIssuesByProject), arg0, arg1)
}

// MockIProjectService is a mock of IProjectService interface.
type MockIProjectService struct {
	ctrl     *gomock.Controller
	recorder *MockIProjectServiceMockRecorder
}

// MockIProjectServiceMockRecorder is the mock recorder for MockIProjectService.
type MockIProjectServiceMockRecorder struct {
	mock *MockIProjectService
}

// NewMockIProjectService creates a new mock instance.
func NewMockIProjectService(ctrl *gomock.Controller) *MockIProjectService {
	mock := &MockIProjectService{ctrl: ctrl}
	mock.recorder = &MockIProjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProjectService) EXPECT() *MockIProjectServiceMockRecorder {
	return m.recorder
}

// GetProjects mocks base method.
func (m *MockIProjectService) GetProjects(arg0 context.Context) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockIProjectServiceMockRecorder) GetProjects(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockIProjectService)(nil).GetProjects), arg0)
}

// GetProjectsByRange mocks base method.
func (m *MockIProjectService) GetProjectsByRange(arg0 context.Context, arg1, arg2 int) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsByRange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsByRange indicates an expected call of GetProjectsByRange.
func (mr *MockIProjectServiceMockRecorder) GetProjectsByRange(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsByRange", reflect.TypeOf((*MockIProjectService)(nil).GetProjectsByRange), arg0, arg1, arg2)
}
