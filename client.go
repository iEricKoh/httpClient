package httpClient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type httpClient struct{ Config }

var client *httpClient

func createDefaultClient() *httpClient {
	return &httpClient{}
}

func Create(config *Config) *httpClient {
	return &httpClient{*config}
}

func Get(url string, options *Options) (*Response, error) {
	return client.Get(url, options)
}

func (h *httpClient) Get(url string, options *Options) (*Response, error) {
	return h.DoRequest("GET", url, options)
}

func (h *httpClient) DoRequest(method, url string, options *Options) (*Response, error) {
	req, err := h.createHttpRequest(method, url, options)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("get response for url=%s got error=%s\n", url, err.Error())
	}

	return &Response{
		Status:        resp.Status,
		StatusCode:    resp.StatusCode,
		Proto:         resp.Proto,
		ProtoMajor:    resp.ProtoMajor,
		ProtoMinor:    resp.ProtoMinor,
		Header:        &resp.Header,
		ContentLength: resp.ContentLength,
		Request:       req,
		Body:          body,
	}, err
}

func (h *httpClient) createHttpRequest(method, url string, options *Options) (*http.Request, error) {
	if options == nil {
		options = &Options{}
	}

	u := h.parseURL(url)

	header := h.populateHeader(options.Header)

	req, err := http.NewRequest(method, u, nil)

	for key, val := range *header {
		req.Header.Add(key, val)
	}

	if options.Query != nil {
		q := req.URL.Query()

		for key, value := range *options.Query {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		req.URL.RawQuery = q.Encode()
	}

	return req, err
}

func (h *httpClient) parseURL(link string) string {
	if h.Host != "" {
		return h.Host + link
	}

	return link
}

func (h *httpClient) populateHeader(header *Header) *Header {
	headers := Header{}

	if h.Header != nil {
		for key, val := range *h.Header {
			headers[key] = val
		}
	}

	if header != nil {
		for key, val := range *header {
			headers[key] = val
		}
	}

	return &headers
}

func init() {
	client = createDefaultClient()
}
