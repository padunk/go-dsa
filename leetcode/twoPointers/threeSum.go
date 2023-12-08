package twopointers

import (
	"math"
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	left, right := 0, len(nums)-1
	for left < right {
		isLookupLeft := math.Abs(float64(nums[left])) < math.Abs(float64(nums[right]))

		if isLookupLeft {
			third := nums[left+1]
			if isValid(nums[left], nums[right], third, result) {
				result = append(result, []int{nums[left], nums[right], third})
			}
			right--
		} else {
			third := nums[right-1]
			if isValid(nums[left], nums[right], third, result) {
				result = append(result, []int{nums[left], nums[right], third})
			}
			left++
		}
	}

	return result
}

func isValid(a, b, c int, n [][]int) bool {
	if a+b+c != 0 {
		return false
	}
	if len(n) < 1 {
		return true
	}

	for i := 0; i < len(n); i++ {
		if slices.Contains(n[i], a) && slices.Contains(n[i], b) && slices.Contains(n[i], c) {
			return false
		}
	}

	return true
}

// ChatGPT
func chatGPThreeSums(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
