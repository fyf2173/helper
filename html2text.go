package helper

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func ExtractTextFromHtml(s string, limit int) string {
	var desc = ""
	var count = 0

	domDocTest := html.NewTokenizer(strings.NewReader(s))
	previousStartTokenTest := domDocTest.Token()

loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			if limit >= 0 && count > limit {
				break
			}
			if len(TxtContent) > 0 {
				if desc == "" {
					desc += fmt.Sprintf("%s ", TxtContent)
				} else {
					desc += fmt.Sprintf(" %s", TxtContent)
				}
				count += len(TxtContent)
			}
		}
	}
	return desc
}
