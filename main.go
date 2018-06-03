package main

import (
	"flag"
	"fmt"
	"github.com/sgoedecke/mashup/markov"
	"github.com/sgoedecke/mashup/scraper"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	numSentencesPtr := flag.String("n", "3", "number of sentences to generate")
	flag.Parse()
	rand.Seed(time.Now().Unix())

	dict := markov.NewMarkov()
	for _, url := range flag.Args() {
		fmt.Println("Processing " + url + "...")
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		text := scraper.ExtractText(resp)
		dict.Add(text)
	}

	numSentences, _ := strconv.Atoi(*numSentencesPtr)
	fmt.Println(dict.GenerateParagraph(numSentences))
}
