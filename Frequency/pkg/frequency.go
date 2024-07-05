package frequency

import (
	"regexp"
	"sort"
	"strings"
)

type WordFreq struct {
	Word  string
	Count int
}

func GetTopWords(text string) []WordFreq {
	stopWords := map[string]struct{}{
		"the": {}, "is": {}, "in": {}, "at": {}, "of": {}, "and": {}, "a": {}, "to": {}, "if": {}, "on": {}, "for": {}, "it": {}, "this": {}, "like": {}, "etc": {},
	}

	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^\w\s]+`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)

	frequency := make(map[string]int)
	for _, word := range words {
		if _, found := stopWords[word]; !found {
			frequency[word]++
		}
	}

	wordFreqs := make([]WordFreq, 0, len(frequency))
	for word, count := range frequency {
		wordFreqs = append(wordFreqs, WordFreq{Word: word, Count: count})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Count > wordFreqs[j].Count
	})

	if len(wordFreqs) > 10 {
		return wordFreqs[:10]
	}
	return wordFreqs
}
