package client

type HttpClient interface {
	Get(url string, v ...interface{}) (HttpReponse, error)
	Post(url string, v ...interface{}) (HttpReponse, error)
}

type HttpReponse interface {
	Bytes() []byte
	ToBytes() ([]byte, error)
	ToJSON(v interface{}) error
}
