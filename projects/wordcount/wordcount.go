package wordcount

import "strings"

func WordCount(text string) map[string]int {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
