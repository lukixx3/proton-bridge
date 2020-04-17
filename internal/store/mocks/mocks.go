// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/internal/store (interfaces: PanicHandler,BridgeUser)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPanicHandler is a mock of PanicHandler interface
type MockPanicHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPanicHandlerMockRecorder
}

// MockPanicHandlerMockRecorder is the mock recorder for MockPanicHandler
type MockPanicHandlerMockRecorder struct {
	mock *MockPanicHandler
}

// NewMockPanicHandler creates a new mock instance
func NewMockPanicHandler(ctrl *gomock.Controller) *MockPanicHandler {
	mock := &MockPanicHandler{ctrl: ctrl}
	mock.recorder = &MockPanicHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPanicHandler) EXPECT() *MockPanicHandlerMockRecorder {
	return m.recorder
}

// HandlePanic mocks base method
func (m *MockPanicHandler) HandlePanic() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandlePanic")
}

// HandlePanic indicates an expected call of HandlePanic
func (mr *MockPanicHandlerMockRecorder) HandlePanic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlePanic", reflect.TypeOf((*MockPanicHandler)(nil).HandlePanic))
}

// MockBridgeUser is a mock of BridgeUser interface
type MockBridgeUser struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeUserMockRecorder
}

// MockBridgeUserMockRecorder is the mock recorder for MockBridgeUser
type MockBridgeUserMockRecorder struct {
	mock *MockBridgeUser
}

// NewMockBridgeUser creates a new mock instance
func NewMockBridgeUser(ctrl *gomock.Controller) *MockBridgeUser {
	mock := &MockBridgeUser{ctrl: ctrl}
	mock.recorder = &MockBridgeUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBridgeUser) EXPECT() *MockBridgeUserMockRecorder {
	return m.recorder
}

// CloseConnection mocks base method
func (m *MockBridgeUser) CloseConnection(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseConnection", arg0)
}

// CloseConnection indicates an expected call of CloseConnection
func (mr *MockBridgeUserMockRecorder) CloseConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockBridgeUser)(nil).CloseConnection), arg0)
}

// GetAddressID mocks base method
func (m *MockBridgeUser) GetAddressID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressID indicates an expected call of GetAddressID
func (mr *MockBridgeUserMockRecorder) GetAddressID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressID", reflect.TypeOf((*MockBridgeUser)(nil).GetAddressID), arg0)
}

// GetPrimaryAddress mocks base method
func (m *MockBridgeUser) GetPrimaryAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrimaryAddress indicates an expected call of GetPrimaryAddress
func (mr *MockBridgeUserMockRecorder) GetPrimaryAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryAddress", reflect.TypeOf((*MockBridgeUser)(nil).GetPrimaryAddress))
}

// GetStoreAddresses mocks base method
func (m *MockBridgeUser) GetStoreAddresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreAddresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStoreAddresses indicates an expected call of GetStoreAddresses
func (mr *MockBridgeUserMockRecorder) GetStoreAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreAddresses", reflect.TypeOf((*MockBridgeUser)(nil).GetStoreAddresses))
}

// ID mocks base method
func (m *MockBridgeUser) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockBridgeUserMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockBridgeUser)(nil).ID))
}

// IsCombinedAddressMode mocks base method
func (m *MockBridgeUser) IsCombinedAddressMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCombinedAddressMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCombinedAddressMode indicates an expected call of IsCombinedAddressMode
func (mr *MockBridgeUserMockRecorder) IsCombinedAddressMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCombinedAddressMode", reflect.TypeOf((*MockBridgeUser)(nil).IsCombinedAddressMode))
}

// IsConnected mocks base method
func (m *MockBridgeUser) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockBridgeUserMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockBridgeUser)(nil).IsConnected))
}

// Logout mocks base method
func (m *MockBridgeUser) Logout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockBridgeUserMockRecorder) Logout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockBridgeUser)(nil).Logout))
}

// UpdateUser mocks base method
func (m *MockBridgeUser) UpdateUser() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockBridgeUserMockRecorder) UpdateUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockBridgeUser)(nil).UpdateUser))
}
