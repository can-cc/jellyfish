// Code generated by MockGen. DO NOT EDIT.
// Source: domain/taco/repository/taco_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	taco "github.com/fwchen/jellyfish/domain/taco"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockRepository) List(userId string, filter taco.TacoFilter) ([]taco.Taco, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userId, filter)
	ret0, _ := ret[0].([]taco.Taco)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List(userId, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), userId, filter)
}

// Save mocks base method
func (m *MockRepository) Save(taco *taco.Taco) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", taco)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockRepositoryMockRecorder) Save(taco interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), taco)
}

// SaveList mocks base method
func (m *MockRepository) SaveList(tacos []taco.Taco) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveList", tacos)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveList indicates an expected call of SaveList
func (mr *MockRepositoryMockRecorder) SaveList(tacos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveList", reflect.TypeOf((*MockRepository)(nil).SaveList), tacos)
}

// FindById mocks base method
func (m *MockRepository) FindById(tacoID string) (*taco.Taco, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", tacoID)
	ret0, _ := ret[0].(*taco.Taco)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockRepositoryMockRecorder) FindById(tacoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), tacoID)
}

// MaxOrderByCreatorId mocks base method
func (m *MockRepository) MaxOrderByCreatorId(userId string) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxOrderByCreatorId", userId)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxOrderByCreatorId indicates an expected call of MaxOrderByCreatorId
func (mr *MockRepositoryMockRecorder) MaxOrderByCreatorId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxOrderByCreatorId", reflect.TypeOf((*MockRepository)(nil).MaxOrderByCreatorId), userId)
}

// MaxOrderByBoxId mocks base method
func (m *MockRepository) MaxOrderByBoxId(userId string) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxOrderByBoxId", userId)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxOrderByBoxId indicates an expected call of MaxOrderByBoxId
func (mr *MockRepositoryMockRecorder) MaxOrderByBoxId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxOrderByBoxId", reflect.TypeOf((*MockRepository)(nil).MaxOrderByBoxId), userId)
}

// Delete mocks base method
func (m *MockRepository) Delete(tacoId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", tacoId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(tacoId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), tacoId)
}
