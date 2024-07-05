package main

import (
	"fmt"

	frequency "github.com/NikKazzzzz/OTUS-PRACTIC/Frequency/pkg"
)

func main() {
	text := `This is a sample text to test more more the frequency analysis function. It includes some common stop words like the, is, and, if, etc. This text has more repetitions of certain words, which should appear in the top results more frequently. For example, the word "text" appears multiple times, as well as "function" and "sample".`
	topWords := frequency.GetTopWords(text)

	fmt.Println("Top 10 most frequent words:")
	for _, wf := range topWords {
		fmt.Printf("%s: %d\n", wf.Word, wf.Count)
	}
}
