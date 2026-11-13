package redo

func Unique(nums []int) []int {
	result := []int{}
	seen := map[int]bool{}

	for _, num := range nums {
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	return result
}
