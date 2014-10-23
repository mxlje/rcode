package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type redirectLogger struct{}

func main() {
	// get URL from cli arguments
	url := os.Args[1]

	// new http client with custom round tripper which
	// itself uses the default transport policy
	client := &http.Client{
		Transport:     redirectLogger{},
		CheckRedirect: redirectCounter}

	// create request
	req, _ := http.NewRequest("GET", url, nil)

	// … with custom user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	printResponse(resp)
}

// RoundTrip only receives a request, executes it and returns the response
// without modifications. If the response code is a redirect of some form
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

// Increases the max redirect count to 20
var redirectCounter = func(req *http.Request, via []*http.Request) error {
	if len(via) >= 20 {
		return errors.New("Stopped after 20 redirects")
	}
	return nil
}

func printResponse(resp *http.Response) {
	fmt.Printf("[%d] %s\n", resp.StatusCode, resp.Request.URL)
}
