package search

func searchInsert(nums []int, target int) int {
	result := -1
	lo := 0
	high := len(nums)

	for lo < high {
		mid := lo + (high-lo)/2
		if target == nums[mid] {
			result = mid
			break
		}

		if target < nums[mid] {
			high = mid
		} else {
			lo = mid + 1
		}
	}

	if result == -1 {
		result = lo
	}
	return result
}
