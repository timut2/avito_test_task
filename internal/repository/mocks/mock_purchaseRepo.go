// Code generated by MockGen. DO NOT EDIT.
// Source: D:/vscodeprojects/golang/avito_test/internal/repository/purchaseRepo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPurchaseRepo is a mock of PurchaseRepo interface.
type MockPurchaseRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseRepoMockRecorder
}

// MockPurchaseRepoMockRecorder is the mock recorder for MockPurchaseRepo.
type MockPurchaseRepoMockRecorder struct {
	mock *MockPurchaseRepo
}

// NewMockPurchaseRepo creates a new mock instance.
func NewMockPurchaseRepo(ctrl *gomock.Controller) *MockPurchaseRepo {
	mock := &MockPurchaseRepo{ctrl: ctrl}
	mock.recorder = &MockPurchaseRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchaseRepo) EXPECT() *MockPurchaseRepoMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockPurchaseRepo) Insert(userID, itemID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", userID, itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockPurchaseRepoMockRecorder) Insert(userID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPurchaseRepo)(nil).Insert), userID, itemID)
}
