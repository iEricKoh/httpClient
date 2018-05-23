package httpClient

import (
	"net/http"
	"time"
)

type Config struct {
	Host          string
	Header        *Header
	Query         *Query
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration
}
