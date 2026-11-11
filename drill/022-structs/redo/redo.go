package redo

import "strings"

func WordCount(s string) map[string]int {
	count := map[string]int{}
	stringToMap := strings.Fields(strings.ToLower(s))

	for _, word := range stringToMap {
		count[word]++
	}

	return count
}
