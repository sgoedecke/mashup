package markov

import (
	"math/rand"
	"regexp"
	"strings"
)

type MarkovDict struct {
	FirstWords []string
	LastWords  []string
	Dict       map[string][]string
}

func NewMarkov() *MarkovDict {
	dict := MarkovDict{}
	dict.Dict = make(map[string][]string)
	return &dict
}

func (m *MarkovDict) Add(s string) {
	re := regexp.MustCompile(`(\s)+`)
	s = re.ReplaceAllString(s, " ")
	s = strings.Replace(s, ";", ". ", -1)
	sentences := strings.Split(s+" ", ". ")
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		m.LastWords = append(m.LastWords, words[len(words)-1])
		for i, word := range words {
			if i == 0 {
				m.FirstWords = append(m.FirstWords, word)
			} else {
				m.Dict[words[i-1]] = append(m.Dict[words[i-1]], word)
			}
		}
	}
}

func (m *MarkovDict) GenerateSentence() string {
	if len(m.FirstWords) == 0 {
		panic("Dictionary is empty!")
	}
	var sentence []string
	finished := false
	sentence = append(sentence, m.FirstWords[rand.Intn(len(m.FirstWords))])
	for !finished {
		last := sentence[len(sentence)-1]
		candidates := m.Dict[last]
		next := candidates[rand.Intn(len(candidates))]
		sentence = append(sentence, next)
		if len(m.Dict[next]) == 0 {
			break
		}
		for _, w := range m.LastWords {
			if next == w && rand.Intn(100) < 2 { // don't always end as soon as possible
				finished = true
				break
			}
		}
	}
	return strings.Join(sentence, " ") + "."
}

func (m *MarkovDict) GenerateParagraph(n int) string {
	paragraph := []string{}
	for i := 0; i < n; i++ {
		paragraph = append(paragraph, m.GenerateSentence())
	}
	return strings.Join(paragraph, " ")
}
