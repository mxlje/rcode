package main

import (
	"net/http"
)

// redirectLogger implements the http.RoundTripper interface
type redirectLogger struct{}

// RoundTrip receives a request, executes it and returns the response without
// any modifications. If the response code is a redirect of some form
// it gets send down the responses channel to be logged
func (l redirectLogger) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := http.DefaultTransport

	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}

	// If the response code is 3xx we print it out
	switch resp.StatusCode {
	case http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusTemporaryRedirect:

		printResponse(resp)
	}
	return
}
