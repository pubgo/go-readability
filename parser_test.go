package readability

import (
	"github.com/pubgo/errors"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
)

func getNodeExcerpt(node *html.Node) string {
	defer errors.Handle()()

	outer := outerHTML(node)
	outer = strings.Join(strings.Fields(outer), " ")
	if len(outer) < 120 {
		return outer
	}
	return outer[:120]
}

func compareArticleContent(result, expected *html.Node) {
	defer errors.Handle()()

	// Make sure number of nodes is same
	resultNodesCount := len(children(result))
	expectedNodesCount := len(children(expected))
	errors.T(resultNodesCount != expectedNodesCount, "number of nodes is different, want %d got %d",
		expectedNodesCount, resultNodesCount)

	resultNode := result
	expectedNode := expected
	for resultNode != nil && expectedNode != nil {
		// Get node excerpt
		resultExcerpt := getNodeExcerpt(resultNode)
		expectedExcerpt := getNodeExcerpt(expectedNode)

		// Compare tag name
		resultTagName := tagName(resultNode)
		expectedTagName := tagName(expectedNode)
		errors.T(resultTagName != expectedTagName, "tag name is different\n"+
			"want    : %s (%s)\n"+
			"got     : %s (%s)",
			expectedTagName, expectedExcerpt,
			resultTagName, resultExcerpt)

		// Compare attributes
		resultAttrCount := len(resultNode.Attr)
		expectedAttrCount := len(expectedNode.Attr)
		errors.T(resultAttrCount != expectedAttrCount, "number of attributes is different\n"+
			"want    : %d (%s)\n"+
			"got     : %d (%s)",
			expectedAttrCount, expectedExcerpt,
			resultAttrCount, resultExcerpt)

		for _, resultAttr := range resultNode.Attr {
			expectedAttrVal := getAttribute(expectedNode, resultAttr.Key)
			switch resultAttr.Key {
			case "href", "src":
				resultAttr.Val = strings.TrimSuffix(resultAttr.Val, "/")
				expectedAttrVal = strings.TrimSuffix(expectedAttrVal, "/")
			}

			errors.T(resultAttr.Val != expectedAttrVal, "attribute %s is different\n"+
				"want    : %s (%s)\n"+
				"got     : %s (%s)",
				resultAttr.Key, expectedAttrVal, expectedExcerpt,
				resultAttr.Val, resultExcerpt)
		}

		// Compare text content
		resultText := strings.TrimSpace(textContent(resultNode))
		expectedText := strings.TrimSpace(textContent(expectedNode))

		resultText = strings.Join(strings.Fields(resultText), " ")
		expectedText = strings.Join(strings.Fields(expectedText), " ")

		comparator := diffmatchpatch.New()
		diffs := comparator.DiffMain(resultText, expectedText, false)

		errors.T(len(diffs) > 1, "text content is different\n"+
			"want  : %s\n"+
			"got   : %s\n"+
			"diffs : %s",
			expectedExcerpt, resultExcerpt,
			comparator.DiffPrettyText(diffs))

		// Move to next node
		ps := Parser{}
		resultNode = ps.getNextNode(resultNode, false)
		expectedNode = ps.getNextNode(expectedNode, false)
	}
}

func TestParser(t *testing.T) {
	errors.Cfg.Debug=true

	defer errors.Debug()

	testDir := "test-pages"
	testItems, err := ioutil.ReadDir(testDir)
	errors.Wrap(err,"failed to read test directory")

	for _, item := range testItems {
		if !item.IsDir() {
			continue
		}

		t.Run(item.Name(), func(t1 *testing.T) {
			defer errors.Debug()

			// Open test file
			testFilePath := fp.Join(testDir, item.Name(), "source.html")
			testFile, err := os.Open(testFilePath)
			if err != nil {
				t1.Errorf("\nfailed to open test file")
			}
			defer errors.Panic(testFile.Close())

			// Open expected result file
			expectedFilePath := fp.Join(testDir, item.Name(), "expected.html")
			expectedFile, err := os.Open(expectedFilePath)
			errors.Wrap(err, "failed to open expected result file")

			defer expectedFile.Close()

			// Parse expected result
			expectedHTML, err := html.Parse(expectedFile)
			errors.Wrap(err, "failed to parse expected result file")

			// Get article from test file
			resultArticle := FromReader(testFile, "http://fakehost/test/page.html")

			// Parse article into HTML
			resultHTML, err := html.Parse(strings.NewReader(resultArticle.Content))
			errors.Wrap(err, "failed to parse test article into HTML")

			// Compare article
			compareArticleContent(resultHTML, expectedHTML)
		})
	}
}

func TestA1(t *testing.T) {

}