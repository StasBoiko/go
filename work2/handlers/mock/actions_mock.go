// Code generated by MockGen. DO NOT EDIT.
// Source: actions.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	models "work2/models"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, task)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServiceMockRecorder) CreateTask(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskService)(nil).CreateTask), ctx, task)
}

// DeleteTask mocks base method.
func (m *MockTaskService) DeleteTask(ctx context.Context, id string) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskServiceMockRecorder) DeleteTask(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskService)(nil).DeleteTask), ctx, id)
}

// FirstTask mocks base method.
func (m *MockTaskService) FirstTask(ctx context.Context, id string, task *models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstTask", ctx, id, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// FirstTask indicates an expected call of FirstTask.
func (mr *MockTaskServiceMockRecorder) FirstTask(ctx, id, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstTask", reflect.TypeOf((*MockTaskService)(nil).FirstTask), ctx, id, task)
}

// GetTasks mocks base method.
func (m *MockTaskService) GetTasks(ctx context.Context, f, v string, userId uint) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", ctx, f, v, userId)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockTaskServiceMockRecorder) GetTasks(ctx, f, v, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockTaskService)(nil).GetTasks), ctx, f, v, userId)
}

// SaveTask mocks base method.
func (m *MockTaskService) SaveTask(ctx context.Context, task models.Task) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTask", ctx, task)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTask indicates an expected call of SaveTask.
func (mr *MockTaskServiceMockRecorder) SaveTask(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTask", reflect.TypeOf((*MockTaskService)(nil).SaveTask), ctx, task)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CheckPassword mocks base method.
func (m *MockUserService) CheckPassword(ctx context.Context, userPass, providedPass string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", ctx, userPass, providedPass)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockUserServiceMockRecorder) CheckPassword(ctx, userPass, providedPass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockUserService)(nil).CheckPassword), ctx, userPass, providedPass)
}

// CreateUserIdTaskId mocks base method.
func (m *MockUserService) CreateUserIdTaskId(ctx context.Context, user models.User, task models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserIdTaskId", ctx, user, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserIdTaskId indicates an expected call of CreateUserIdTaskId.
func (mr *MockUserServiceMockRecorder) CreateUserIdTaskId(ctx, user, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserIdTaskId", reflect.TypeOf((*MockUserService)(nil).CreateUserIdTaskId), ctx, user, task)
}

// CreateUserRecord mocks base method.
func (m *MockUserService) CreateUserRecord(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRecord", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserRecord indicates an expected call of CreateUserRecord.
func (mr *MockUserServiceMockRecorder) CreateUserRecord(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRecord", reflect.TypeOf((*MockUserService)(nil).CreateUserRecord), ctx, user)
}

// GetFirstUser mocks base method.
func (m *MockUserService) GetFirstUser(ctx context.Context, email, token string, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstUser", ctx, email, token, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstUser indicates an expected call of GetFirstUser.
func (mr *MockUserServiceMockRecorder) GetFirstUser(ctx, email, token, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstUser", reflect.TypeOf((*MockUserService)(nil).GetFirstUser), ctx, email, token, user)
}

// GetUser mocks base method.
func (m *MockUserService) GetUser(ctx context.Context, email string, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceMockRecorder) GetUser(ctx, email, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), ctx, email, user)
}

// HashPassword mocks base method.
func (m *MockUserService) HashPassword(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockUserServiceMockRecorder) HashPassword(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockUserService)(nil).HashPassword), ctx, user)
}

// SaveUserToken mocks base method.
func (m *MockUserService) SaveUserToken(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserToken", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveUserToken indicates an expected call of SaveUserToken.
func (mr *MockUserServiceMockRecorder) SaveUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserToken", reflect.TypeOf((*MockUserService)(nil).SaveUserToken), ctx, user)
}
