package battlenet

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// Battle.net client being tested.
	client *Client

	// Test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("us", nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
