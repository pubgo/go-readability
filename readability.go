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
	"bytes"
	"github.com/pubgo/assert"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FromReader parses input from an `io.Reader` and returns the
// readable content. It's the wrapper for `Parser.Parse()` and useful
// if you only want to use the default parser.
func FromReader(input io.Reader, pageURL string) Article {
	return NewParser().Parse(input, pageURL)
}

// IsReadable decides whether or not the document is reader-able
// without parsing the whole thing. It's the wrapper for
// `Parser.IsReadable()` and useful if you only use the default parser.
func IsReadable(input io.Reader) bool {
	parser := NewParser()
	return parser.IsReadable(input)
}

// FromURL fetch the web page from specified url, check if it's
// readable, then parses the response to find the readable content.
func FromURL(pageURL string, timeout time.Duration) Article {
	defer assert.Panic("FromURL Error")

	// Make sure URL is valid
	_, err := url.ParseRequestURI(pageURL)
	assert.ErrWrap(err, "failed to parse URL")

	// Fetch page from URL
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(pageURL)
	assert.ErrWrap(err, "failed to fetch the page")

	defer assert.Throw(resp.Body.Close())

	// Make sure content type is HTML
	cp := resp.Header.Get("Content-Type")
	assert.T(!strings.Contains(cp, "text/html"), "URL is not a HTML document")

	// Check if the page is readable
	var buffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &buffer)

	parser := NewParser()
	assert.T(!parser.IsReadable(tee), "the page is not readable")

	// Parse content
	return parser.Parse(&buffer, pageURL)
}
