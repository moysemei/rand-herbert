package redo

func ReverseInPlace(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i

		nums[i], nums[j] = nums[j], nums[i]
	}
}
