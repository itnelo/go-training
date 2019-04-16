package exercises

import (
	"strings"

	"golang.org/x/tour/wc"
)

func mapWordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordOccurrenceMap := make(map[string]int)

	for _, word := range words {
		wordOccurrenceMap[word]++
	}

	return wordOccurrenceMap
}

func wordCount() {
	wc.Test(mapWordCount)
}
