// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prometheus/client_golang/api (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockClient) Do(arg0 context.Context, arg1 *http.Request) (*http.Response, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Do indicates an expected call of Do
func (mr *MockClientMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClient)(nil).Do), arg0, arg1)
}

// URL mocks base method
func (m *MockClient) URL(arg0 string, arg1 map[string]string) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL", arg0, arg1)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockClientMockRecorder) URL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockClient)(nil).URL), arg0, arg1)
}
