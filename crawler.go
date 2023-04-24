package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return findRecurse(doc.Find("body").Children(), "p, h1, h2, h3, h4, h5, h6, ul, ol, pre, blockquote"), nil
}

const newLine = "\r\n"

func findRecurse(children *goquery.Selection, tags string) string {
	if children == nil {
		return ""
	}
	res := ""
	children.Each(func(i int, s *goquery.Selection) {
		if s.Is(tags) {
			str := strings.TrimSpace(s.Text())
			if str != "" {
				res += str + newLine
			}
		} else {
			res += findRecurse(s.Children(), tags)
		}
	})
	return res
}
