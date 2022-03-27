package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)

	for _, str := range(fields) {
		_, ok := result[str]

		if ok == true {
			result[str]++
		} else {
			result[str] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
