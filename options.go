package httpClient

import "net/http/cookiejar"

type Header map[string]string
type Query map[string]interface{}
type Form map[string]string

type Options struct {
	Header *Header
	Query  *Query
	Form   *Form
	Jar    *cookiejar.Jar
}
