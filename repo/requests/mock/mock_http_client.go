// Code generated by MockGen. DO NOT EDIT.
// Source: requests.go
//
// Generated by this command:
//
//	mockgen -destination=mock/mock_http_client.go -package mock -source=requests.go
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHttpClientInterface is a mock of HttpClientInterface interface.
type MockHttpClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientInterfaceMockRecorder
}

// MockHttpClientInterfaceMockRecorder is the mock recorder for MockHttpClientInterface.
type MockHttpClientInterfaceMockRecorder struct {
	mock *MockHttpClientInterface
}

// NewMockHttpClientInterface creates a new mock instance.
func NewMockHttpClientInterface(ctrl *gomock.Controller) *MockHttpClientInterface {
	mock := &MockHttpClientInterface{ctrl: ctrl}
	mock.recorder = &MockHttpClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClientInterface) EXPECT() *MockHttpClientInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockHttpClientInterface) Delete(url string, headers map[string]string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", url, headers)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockHttpClientInterfaceMockRecorder) Delete(url, headers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHttpClientInterface)(nil).Delete), url, headers)
}

// Get mocks base method.
func (m *MockHttpClientInterface) Get(url string, headers map[string]string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url, headers)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientInterfaceMockRecorder) Get(url, headers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClientInterface)(nil).Get), url, headers)
}

// Post mocks base method.
func (m *MockHttpClientInterface) Post(url string, headers map[string]string, postBody []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", url, headers, postBody)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHttpClientInterfaceMockRecorder) Post(url, headers, postBody any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClientInterface)(nil).Post), url, headers, postBody)
}
