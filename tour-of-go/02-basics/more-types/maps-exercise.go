package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	arr := strings.Fields(s)
	for _, val := range arr {
		count := wordCount[val]
		wordCount[val] = count + 1
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
