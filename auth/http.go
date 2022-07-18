package auth

import (
	"context"
)

type HTTPClient interface {
	NewRequest() HTTPRequest
}

type HTTPRequest interface {
	SetFormData(data map[string]string) HTTPRequest
	SetResult(result interface{}) HTTPRequest
	SetError(error interface{}) HTTPRequest
	SetHeaders(hdrs map[string]string) HTTPRequest
	Get(url string) (HTTPResponse, error)
	Post(url string) (HTTPResponse, error)
	Put(url string) (HTTPResponse, error)
	SetContext(ctx context.Context) HTTPRequest
	DisableTrace() HTTPRequest
	EnableTrace() HTTPRequest
}

type HTTPResponse interface {
	IsSuccess() bool
	IsError() bool
	Result() interface{}
	Error() interface{}
	UnmarshalJson(v interface{}) error
	Bytes() []byte
	String() string
	ToString() (string, error)
	ToBytes() ([]byte, error)
	GetStatus() string
	GetStatusCode() int
	GetHeader(key string) string
	GetHeaderValues(key string) []string
}
