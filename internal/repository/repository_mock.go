// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nurhidaylma/parking-app/internal/model"
)

// MockRepoInterface is a mock of RepoInterface interface.
type MockRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepoInterfaceMockRecorder
}

// MockRepoInterfaceMockRecorder is the mock recorder for MockRepoInterface.
type MockRepoInterfaceMockRecorder struct {
	mock *MockRepoInterface
}

// NewMockRepoInterface creates a new mock instance.
func NewMockRepoInterface(ctrl *gomock.Controller) *MockRepoInterface {
	mock := &MockRepoInterface{ctrl: ctrl}
	mock.recorder = &MockRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoInterface) EXPECT() *MockRepoInterfaceMockRecorder {
	return m.recorder
}

// ReadParkingSpots mocks base method.
func (m *MockRepoInterface) ReadParkingSpots() ([]model.ParkingSpot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadParkingSpots")
	ret0, _ := ret[0].([]model.ParkingSpot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadParkingSpots indicates an expected call of ReadParkingSpots.
func (mr *MockRepoInterfaceMockRecorder) ReadParkingSpots() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadParkingSpots", reflect.TypeOf((*MockRepoInterface)(nil).ReadParkingSpots))
}

// ReadVehicles mocks base method.
func (m *MockRepoInterface) ReadVehicles() ([]model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadVehicles")
	ret0, _ := ret[0].([]model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadVehicles indicates an expected call of ReadVehicles.
func (mr *MockRepoInterfaceMockRecorder) ReadVehicles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadVehicles", reflect.TypeOf((*MockRepoInterface)(nil).ReadVehicles))
}

// WriteInitialVehicle mocks base method.
func (m *MockRepoInterface) WriteInitialVehicle() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteInitialVehicle")
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteInitialVehicle indicates an expected call of WriteInitialVehicle.
func (mr *MockRepoInterfaceMockRecorder) WriteInitialVehicle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteInitialVehicle", reflect.TypeOf((*MockRepoInterface)(nil).WriteInitialVehicle))
}

// WriteParkingSpots mocks base method.
func (m *MockRepoInterface) WriteParkingSpots(arg0 []model.ParkingSpot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteParkingSpots", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteParkingSpots indicates an expected call of WriteParkingSpots.
func (mr *MockRepoInterfaceMockRecorder) WriteParkingSpots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteParkingSpots", reflect.TypeOf((*MockRepoInterface)(nil).WriteParkingSpots), arg0)
}

// WriteVehicles mocks base method.
func (m *MockRepoInterface) WriteVehicles(arg0 []model.Vehicle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteVehicles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteVehicles indicates an expected call of WriteVehicles.
func (mr *MockRepoInterfaceMockRecorder) WriteVehicles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteVehicles", reflect.TypeOf((*MockRepoInterface)(nil).WriteVehicles), arg0)
}
