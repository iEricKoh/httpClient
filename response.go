package httpClient

import "net/http"

type Response struct {
	Status        string
	StatusCode    int
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	Header        *http.Header
	ContentLength int64
	Request       *http.Request
	Response      *http.Response
	Body          []byte
}
