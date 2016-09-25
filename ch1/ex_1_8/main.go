// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	for _, u := range os.Args[1:] {
		parsedURL, err := url.Parse(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		if parsedURL.Scheme == "" {
			parsedURL.Scheme = "http"
		}

		resp, err := http.Get(parsedURL.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", parsedURL.String(), err)
			os.Exit(1)
		}
	}
}
