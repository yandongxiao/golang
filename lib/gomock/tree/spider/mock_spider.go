// Code generated by MockGen. DO NOT EDIT.
// Source: tree/spider (interfaces: Spider)

// Package spider is a generated GoMock package.
package spider

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSpider is a mock of Spider interface
type MockSpider struct {
	ctrl     *gomock.Controller
	recorder *MockSpiderMockRecorder
}

// MockSpiderMockRecorder is the mock recorder for MockSpider
type MockSpiderMockRecorder struct {
	mock *MockSpider
}

// NewMockSpider creates a new mock instance
func NewMockSpider(ctrl *gomock.Controller) *MockSpider {
	mock := &MockSpider{ctrl: ctrl}
	mock.recorder = &MockSpiderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpider) EXPECT() *MockSpiderMockRecorder {
	return m.recorder
}

// GetBody mocks base method
func (m *MockSpider) GetBody() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBody indicates an expected call of GetBody
func (mr *MockSpiderMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockSpider)(nil).GetBody))
}
