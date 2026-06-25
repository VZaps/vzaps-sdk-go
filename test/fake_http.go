package test

import (
	"io"
	"net/http"
	"strings"
)

type fakeResponse struct {
	status int
	body   string
}

type fakeHTTP struct {
	requests  []*http.Request
	bodies    [][]byte
	responses []fakeResponse
}

func (f *fakeHTTP) Do(req *http.Request) (*http.Response, error) {
	body, _ := io.ReadAll(req.Body)
	f.requests = append(f.requests, req)
	f.bodies = append(f.bodies, body)
	index := len(f.requests) - 1
	response := f.responses[index]
	if response.status == 0 {
		response.status = http.StatusOK
	}
	return &http.Response{
		StatusCode: response.status,
		Status:     http.StatusText(response.status),
		Body:       io.NopCloser(strings.NewReader(response.body)),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}, nil
}
