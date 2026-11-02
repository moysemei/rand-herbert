package redo

import (
	"strconv"
	"strings"
)

func SumNumbers(s string) (int, bool) {
	sum := 0

	for _, word := range strings.Fields(s) {
		number, err := strconv.Atoi(word)

		if err != nil {
			return 0, false
		}

		sum += number
	}

	return sum, true
}
