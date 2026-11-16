package redo

import "maps"

func charfrequency(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}

	return counts
}

func IsAnagram(a, b string) bool {
	frequencyA := charfrequency(a)
	frequencyB := charfrequency(b)

	return maps.Equal(frequencyA, frequencyB)
}
