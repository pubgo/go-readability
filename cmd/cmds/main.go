package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/errors"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pubgo/go-readability"
	"github.com/spf13/cobra"
)

func initGoReadabilityCmd(cmd *cobra.Command) *cobra.Command {
	return cmd
}

var GoReadabilityCmd = initGoReadabilityCmd(&cobra.Command{
	Use:   "go-readability [flags] source",
	Args:  cobra.ExactArgs(1),
	Run:   rootCmdHandler,
	Short: "go-readability is parser to fetch readable content of a web page",
	Long: "go-readability is parser to fetch the readable content of a web page.\n" +
		"The source can be an url or an existing file in your storage.",
})

func rootCmdHandler(cmd *cobra.Command, args []string) {
	defer errors.Handle(func() {})

	// Get cmd parameter
	srcPath := args[0]
	metadataOnly, _ := cmd.Flags().GetBool("metadata")

	// Open or fetch web page that will be parsed
	var (
		pageURL string
		resp    *http.Response
	)

	if isURL(srcPath) {
		resp, err := http.Get(srcPath)
		errors.Wrap(err, "failed to fetch web page")
		defer errors.Panic(resp.Body.Close())

		pageURL = srcPath
	} else {
		srcFile, err := os.Open(srcPath)
		errors.Wrap(err, "failed to open source file")
		defer errors.Panic(srcFile.Close())

		pageURL = "http://fakehost.com"
	}

	// Get readable content from the reader
	article := readability.FromReader(resp.Body, pageURL)

	// Print the article (or its metadata) to stdout
	if metadataOnly {
		metadata := map[string]interface{}{
			"title":   article.Title,
			"byline":  article.Byline,
			"excerpt": article.Excerpt,
			"image":   article.Image,
			"favicon": article.Favicon,
		}

		prettyJSON, err := json.MarshalIndent(&metadata, "", "    ")
		errors.Wrap(err, "failed to write metadata file")

		fmt.Println(string(prettyJSON))
		return
	}

	fmt.Println(article.Content)
}

func isURL(path string) bool {
	_url, err := url.ParseRequestURI(path)
	return err == nil && strings.HasPrefix(_url.Scheme, "http")
}
