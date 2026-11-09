package redo

import "strings"

func LongestWord(s string) string {
	biggestWord := ""
	cutMyPhrase := strings.Fields(s)

	for _, word := range cutMyPhrase {
		if len(word) > len(biggestWord) {
			biggestWord = word
		}
	}

	return biggestWord
}
