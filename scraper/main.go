package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

// extract only the text from a http response
func ExtractText(resp *http.Response) string {
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	s := string(text)

	p := strings.NewReader(s)
	doc, _ := goquery.NewDocumentFromReader(p)

	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	doc.Find("styles").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("head").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	return doc.Text()
}
