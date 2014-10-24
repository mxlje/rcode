package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	userAgentChrome  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36"
	maxRedirectCount = 20
)

var (
	url                 string
	csvOutput           bool
	validHTTPScheme, _  = regexp.Compile("^https?:\\/\\/")
	errTooManyRedirects = errors.New(fmt.Sprintf("Stopped after %d redirects.", maxRedirectCount))
)

func init() {
	// get URL from cli arguments and add scheme if not present
	url = os.Args[1]
	if valid := validHTTPScheme.MatchString(url); !valid {
		url = "http://" + url
	}

	// TODO: use a real flag parser
	if len(os.Args) == 3 && os.Args[2] == "--csv" {
		csvOutput = true
	}
}

func main() {
	// new http client with custom round tripper which
	// itself uses the default transport policy
	client := &http.Client{
		Transport:     redirectLogger{},
		CheckRedirect: redirectPolicyFunc,
	}

	// create request with a custom user agent
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgentChrome)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *http.Response) {
	if csvOutput {
		fmt.Printf("%d,%s\n", resp.StatusCode, resp.Request.URL)
	} else {
		fmt.Printf("[%d] %s\n", resp.StatusCode, resp.Request.URL)
	}
}
