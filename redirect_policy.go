package main

import (
	"net/http"
)

var redirectPolicyFunc = func(req *http.Request, via []*http.Request) error {
	if len(via) >= maxRedirectCount {
		return errTooManyRedirects
	}
	return nil
}
