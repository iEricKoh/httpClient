package httpClient

type Header map[string]string
type Query map[string]interface{}

type Options struct {
	Header *Header
	Query  *Query
}
