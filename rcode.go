package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type LogRedirects struct {
	Transport http.RoundTripper
}

func main() {
	// get URL from cli arguments
	url := os.Args[1]

	// new http client with custom round tripper
	client := &http.Client{Transport: LogRedirects{}}

	// prepare request
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)

	// … with custom user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[%d] %s\n", resp.StatusCode, resp.Request.URL)
}

// RoundTrip only receives a request and returns the response without modifying it
// Basically we’re doing a regular request and if the code status code is 3xx then we print it out
func (l LogRedirects) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport

	if t == nil {
		t = http.DefaultTransport
	}

	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusTemporaryRedirect:

		// other way of displaying it
		// newLoc, _ := resp.Location()
		// fmt.Printf("[%d] %s\n", resp.StatusCode, newLoc.String())

		fmt.Printf("[%d] %s\n", resp.StatusCode, req.URL)
	}

	return
}
