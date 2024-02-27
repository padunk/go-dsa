package search

// func searchRange(nums []int, target int) []int {

// }

func searchRangeGPT(nums []int, target int) []int {
	left := binarySearchGPT(nums, target, true)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := binarySearchGPT(nums, target, false) - 1
	return []int{left, right}
}

func binarySearchGPT(nums []int, target int, leftBias bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target || (leftBias && nums[mid] == target) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
