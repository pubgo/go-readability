package readability

import (
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
	"task_test/3d/assert"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
)

func getNodeExcerpt(node *html.Node) string {
	outer := outerHTML(node)
	outer = strings.Join(strings.Fields(outer), " ")
	if len(outer) < 120 {
		return outer
	}
	return outer[:120]
}

func compareArticleContent(result, expected *html.Node) {
	// Make sure number of nodes is same
	resultNodesCount := len(children(result))
	expectedNodesCount := len(children(expected))
	assert.T(resultNodesCount != expectedNodesCount, "number of nodes is different, want %d got %d",
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
		assert.T(resultTagName != expectedTagName, "tag name is different\n"+
			"want    : %s (%s)\n"+
			"got     : %s (%s)",
			expectedTagName, expectedExcerpt,
			resultTagName, resultExcerpt)

		// Compare attributes
		resultAttrCount := len(resultNode.Attr)
		expectedAttrCount := len(expectedNode.Attr)
		assert.T(resultAttrCount != expectedAttrCount, "number of attributes is different\n"+
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

			assert.T(resultAttr.Val != expectedAttrVal, "attribute %s is different\n"+
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

		assert.T(len(diffs) > 1, "text content is different\n"+
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

func Test_parser(t *testing.T) {
	testDir := "test-pages"
	testItems, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Errorf("\nfailed to read test directory")
	}

	for _, item := range testItems {
		if !item.IsDir() {
			continue
		}

		t.Run(item.Name(), func(t1 *testing.T) {
			// Open test file
			testFilePath := fp.Join(testDir, item.Name(), "source.html")
			testFile, err := os.Open(testFilePath)
			if err != nil {
				t1.Errorf("\nfailed to open test file")
			}
			defer assert.Throw(testFile.Close())

			// Open expected result file
			expectedFilePath := fp.Join(testDir, item.Name(), "expected.html")
			expectedFile, err := os.Open(expectedFilePath)
			assert.ErrWrap(err, "failed to open expected result file")

			defer assert.Throw(expectedFile.Close())

			// Parse expected result
			expectedHTML, err := html.Parse(expectedFile)
			assert.ErrWrap(err, "failed to parse expected result file")

			// Get article from test file
			resultArticle := FromReader(testFile, "http://fakehost/test/page.html")

			// Parse article into HTML
			resultHTML, err := html.Parse(strings.NewReader(resultArticle.Content))
			assert.ErrWrap(err, "failed to parse test article into HTML")

			// Compare article
			compareArticleContent(resultHTML, expectedHTML)
		})
	}
}
