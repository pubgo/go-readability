package readability

import (
	"github.com/PuerkitoBio/purell"
	"github.com/pubgo/errors"
	"log"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// indexOf returns the position of the first occurrence of a
// specified  value in a string array. Returns -1 if the
// value to search for never occurs.
func indexOf(array []string, key string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return i
		}
	}
	return -1
}

// wordCount returns number of word in str.
func wordCount(str string) int {
	return len(strings.Fields(str))
}

// renderToFile ender an element and save it to file.
// It will panic if it fails to create destination file.
func renderToFile(element *html.Node, filename string) {
	defer errors.Handle()()

	dstFile, err := os.Create(filename)
	errors.Wrap(err, "failed to create file")

	defer dstFile.Close()

	errors.Panic(html.Render(dstFile, element))
}

func ToAbsoluteURI(u string, base *url.URL) string {
	defer errors.Handle()()

	u = strings.ReplaceAll(strings.TrimSpace(u), " ", "")

	if u == "" {
		if Cfg.Debug {
			log.Println("url is null")
		}

		return ""
	}

	// If it is hash tag, return as it is
	if u[:1] == "#" {
		return u
	}

	if strings.HasPrefix(u, "data") {
		if Cfg.Debug {
			log.Println("url HasPrefix data")
		}
		return ""
	}

	u = errors.If(strings.HasPrefix(u, "//"), base.Scheme+":"+u, u).(string)
	u = errors.If(strings.HasPrefix(u, "/"), base.Scheme+"://"+base.Host+u, u).(string)

	// If it is already an absolute URL, return as it is
	_url, err := url.ParseRequestURI(u)
	errors.Wrap(err, "url ParseRequestURI error")
	if err == nil && _url.Scheme != "" && _url.Hostname() != "" {
		return u
	}

	_u, err := purell.NormalizeURLString(
		base.ResolveReference(_url).String(),
		purell.FlagsUsuallySafeGreedy|purell.FlagRemoveDuplicateSlashes|purell.FlagRemoveFragment,
	)
	errors.Wrap(err, "purell NormalizeURLString error")
	return _u
}
