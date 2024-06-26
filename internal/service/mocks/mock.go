// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockURL is a mock of URL interface.
type MockURL struct {
	ctrl     *gomock.Controller
	recorder *MockURLMockRecorder
}

// MockURLMockRecorder is the mock recorder for MockURL.
type MockURLMockRecorder struct {
	mock *MockURL
}

// NewMockURL creates a new mock instance.
func NewMockURL(ctrl *gomock.Controller) *MockURL {
	mock := &MockURL{ctrl: ctrl}
	mock.recorder = &MockURLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURL) EXPECT() *MockURLMockRecorder {
	return m.recorder
}

// CreateShortURL mocks base method.
func (m *MockURL) CreateShortURL(originalURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortURL", originalURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShortURL indicates an expected call of CreateShortURL.
func (mr *MockURLMockRecorder) CreateShortURL(originalURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortURL", reflect.TypeOf((*MockURL)(nil).CreateShortURL), originalURL)
}

// GetOriginalURL mocks base method.
func (m *MockURL) GetOriginalURL(shortURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURL", shortURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURL indicates an expected call of GetOriginalURL.
func (mr *MockURLMockRecorder) GetOriginalURL(shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURL", reflect.TypeOf((*MockURL)(nil).GetOriginalURL), shortURL)
}
