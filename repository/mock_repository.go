// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository.go

// Package repository is a generated GoMock package.
package repository

import (
	model "casestudy/backend-repo/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockRepository) CreateTodo(todo model.SendTodoElements) (model.SendTodoElements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", todo)
	ret0, _ := ret[0].(model.SendTodoElements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockRepositoryMockRecorder) CreateTodo(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockRepository)(nil).CreateTodo), todo)
}

// DeleteAll mocks base method.
func (m *MockRepository) DeleteAll(theList []model.TodoElements) ([]model.TodoElements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", theList)
	ret0, _ := ret[0].([]model.TodoElements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRepositoryMockRecorder) DeleteAll(theList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRepository)(nil).DeleteAll), theList)
}

// GetTodoElements mocks base method.
func (m *MockRepository) GetTodoElements() ([]model.TodoElements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoElements")
	ret0, _ := ret[0].([]model.TodoElements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoElements indicates an expected call of GetTodoElements.
func (mr *MockRepositoryMockRecorder) GetTodoElements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoElements", reflect.TypeOf((*MockRepository)(nil).GetTodoElements))
}
