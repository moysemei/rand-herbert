package redo

// Reverse returns s with its runes in reverse order.
func Reverse(s string) string {
	result := ""

	for _, r := range s {
		result = string(r) + result
	}

	return result
}
