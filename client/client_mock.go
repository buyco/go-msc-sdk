// Code generated by MockGen. DO NOT EDIT.
// Source: ./client/http.go

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpClient is a mock of HttpClient interface.
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientMockRecorder
}

// MockHttpClientMockRecorder is the mock recorder for MockHttpClient.
type MockHttpClientMockRecorder struct {
	mock *MockHttpClient
}

// NewMockHttpClient creates a new mock instance.
func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &MockHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClient) EXPECT() *MockHttpClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHttpClient) Get(url string, v ...interface{}) (HttpReponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(HttpReponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientMockRecorder) Get(url interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClient)(nil).Get), varargs...)
}

// Post mocks base method.
func (m *MockHttpClient) Post(url string, v ...interface{}) (HttpReponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].(HttpReponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHttpClientMockRecorder) Post(url interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClient)(nil).Post), varargs...)
}

// MockHttpReponse is a mock of HttpReponse interface.
type MockHttpReponse struct {
	ctrl     *gomock.Controller
	recorder *MockHttpReponseMockRecorder
}

// MockHttpReponseMockRecorder is the mock recorder for MockHttpReponse.
type MockHttpReponseMockRecorder struct {
	mock *MockHttpReponse
}

// NewMockHttpReponse creates a new mock instance.
func NewMockHttpReponse(ctrl *gomock.Controller) *MockHttpReponse {
	mock := &MockHttpReponse{ctrl: ctrl}
	mock.recorder = &MockHttpReponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpReponse) EXPECT() *MockHttpReponseMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockHttpReponse) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockHttpReponseMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockHttpReponse)(nil).Bytes))
}

// ToBytes mocks base method.
func (m *MockHttpReponse) ToBytes() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToBytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToBytes indicates an expected call of ToBytes.
func (mr *MockHttpReponseMockRecorder) ToBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToBytes", reflect.TypeOf((*MockHttpReponse)(nil).ToBytes))
}

// ToJSON mocks base method.
func (m *MockHttpReponse) ToJSON(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToJSON", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToJSON indicates an expected call of ToJSON.
func (mr *MockHttpReponseMockRecorder) ToJSON(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToJSON", reflect.TypeOf((*MockHttpReponse)(nil).ToJSON), v)
}
