// Code generated by MockGen. DO NOT EDIT.
// Source: ../transfer.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransferer is a mock of Transferer interface
type MockTransferer struct {
	ctrl     *gomock.Controller
	recorder *MockTransfererMockRecorder
}

// MockTransfererMockRecorder is the mock recorder for MockTransferer
type MockTransfererMockRecorder struct {
	mock *MockTransferer
}

// NewMockTransferer creates a new mock instance
func NewMockTransferer(ctrl *gomock.Controller) *MockTransferer {
	mock := &MockTransferer{ctrl: ctrl}
	mock.recorder = &MockTransfererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransferer) EXPECT() *MockTransfererMockRecorder {
	return m.recorder
}

// TransferFunds mocks base method
func (m *MockTransferer) TransferFunds(ctx context.Context, from, to string, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFunds", ctx, from, to, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferFunds indicates an expected call of TransferFunds
func (mr *MockTransfererMockRecorder) TransferFunds(ctx, from, to, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFunds", reflect.TypeOf((*MockTransferer)(nil).TransferFunds), ctx, from, to, amount)
}
