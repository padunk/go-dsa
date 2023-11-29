package leetcode

func productExceptSelf(nums []int) []int {
	length := len(nums)
	leftProducts := make([]int, length)
	rightProducts := make([]int, length)
	result := make([]int, length)

	// calculate left products
	left := 1
	for i := 0; i < length; i++ {
		leftProducts[i] = left
		left *= nums[i]
	}

	// calculate right products
	right := 1
	for i := length - 1; i >= 0; i-- {
		rightProducts[i] = right
		right *= nums[i]
	}

	for i := 0; i < length; i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}
