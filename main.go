package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sgoedecke/mashup/markov"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	numSentencesPtr := flag.String("n", "3", "number of sentences to generate")
	flag.Parse()
	rand.Seed(time.Now().Unix())

	dict := NewMarkov()
	fmt.Println(flag.Args())
	for _, url := range flag.Args() {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		text := ExtractText(resp)
		dict.Add(text)
	}

	numSentences, _ := strconv.Atoi(*numSentencesPtr)
	fmt.Println(dict.GenerateParagraph(numSentences))
}

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
