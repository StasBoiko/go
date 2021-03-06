// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	models "work2/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, user)
}

// CreateUserIdTaskId mocks base method.
func (m *MockUserRepository) CreateUserIdTaskId(ctx context.Context, user *models.User, task *models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserIdTaskId", ctx, user, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserIdTaskId indicates an expected call of CreateUserIdTaskId.
func (mr *MockUserRepositoryMockRecorder) CreateUserIdTaskId(ctx, user, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserIdTaskId", reflect.TypeOf((*MockUserRepository)(nil).CreateUserIdTaskId), ctx, user, task)
}

// FirstUser mocks base method.
func (m *MockUserRepository) FirstUser(ctx context.Context, user *models.User, email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstUser", ctx, user, email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// FirstUser indicates an expected call of FirstUser.
func (mr *MockUserRepositoryMockRecorder) FirstUser(ctx, user, email, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstUser", reflect.TypeOf((*MockUserRepository)(nil).FirstUser), ctx, user, email, token)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(ctx context.Context, user *models.User, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, user, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(ctx, user, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), ctx, user, email)
}

// SaveUser mocks base method.
func (m *MockUserRepository) SaveUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserRepositoryMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserRepository)(nil).SaveUser), ctx, user)
}
