package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordCountMap[words[i]] = wordCountMap[words[i]] + 1
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
