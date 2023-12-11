package window

import (
	"math"
	"slices"
)

// with sorted
func SortedMinSubArrayLen(target int, nums []int) int {
	slices.Sort(nums)
	sum := 0
	count := 0

	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		count++
		if sum >= target {
			return count
		}

	}

	if sum != 0 {
		return 0
	}

	return count
}

// ChatGPT
func GPTMinSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt64
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt64 {
		return 0
	}

	return minLength
}
