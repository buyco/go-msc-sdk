package auth

import "net/http"

type HttpClient interface {
	Get(url string, v ...interface{}) (HttpReponse, error)
	Post(url string, v ...interface{}) (HttpReponse, error)
	Put(url string, v ...interface{}) (HttpReponse, error)
	Delete(url string, v ...interface{}) (HttpReponse, error)
}

type HttpReponse interface {
	Request() *http.Request
	Response() *http.Response
	Bytes() []byte
	ToBytes() ([]byte, error)
	ToJSON(v interface{}) error
}
