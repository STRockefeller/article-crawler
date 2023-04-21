package crawler

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return FindRecurse(doc.Find("body").Children(), "p, h1, h2, h3, h4, h5, h6, ul, ol, pre, blockquote")
}

const newLine = "\r\n"

func FindRecurse(children *goquery.Selection, tags string) string {
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
			res += FindRecurse(s.Children(), tags)
		}
	})
	return res
}
