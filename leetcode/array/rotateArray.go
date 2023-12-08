package leetcode

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	for k > 0 {
		lastValue := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = lastValue

		k--
	}
}

func rotateChatGPT(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
