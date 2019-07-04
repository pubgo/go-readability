// Package readability is a Go package that find the main readable
// content from a HTML page. It works by removing clutter like buttons,
// ads, background images, script, etc.
//
// This package is based from Readability.js by Mozilla, and written line
// by line to make sure it looks and works as similar as possible. This
// way, hopefully all web page that can be parsed by Readability.js
// are parse-able by go-readability as well.
package readability

import (
	"github.com/pubgo/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FromReader parses input from an `io.Reader` and returns the
// readable content. It's the wrapper for `Parser.Parse()` and useful
// if you only want to use the default parser.
func FromReader(input io.Reader, pageURL string) (art *Article) {
	defer errors.Handle()()

	p := NewParser()
	art = p.Parse(input, pageURL)
	p.free()

	return
}

// FromURL fetch the web page from specified url, check if it's
// readable, then parses the response to find the readable content.
func FromURL(pageURL string, timeout time.Duration) (art *Article) {
	defer errors.Handle()()

	// Make sure URL is valid
	_, err := url.ParseRequestURI(pageURL)
	errors.Wrap(err, "failed to parse URL(%s)", pageURL)

	// Fetch page from URL
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(pageURL)
	errors.Wrap(err, "failed to fetch the page")
	defer resp.Body.Close()

	// Make sure content type is HTML
	cp := resp.Header.Get("Content-Type")
	errors.T(!strings.Contains(cp, "text/html"), "URL is not a HTML document")

	// Parse content
	parser := NewParser()
	art = parser.Parse(resp.Body, pageURL)
	parser.free()

	return
}
