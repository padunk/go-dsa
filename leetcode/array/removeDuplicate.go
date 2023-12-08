package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	k, p := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if k != nums[i] {
			k = nums[i]
			nums[p] = nums[i]
			p++
		}
	}

	return p
}

func removeDuplicatesChatGPT(nums []int) int {
	// Edge case: empty array
	if len(nums) == 0 {
		return 0
	}

	// Initialize a pointer for placing unique elements
	k := 1

	// Iterate through the input array starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the current element is different from the previous one
		if nums[i] != nums[i-1] {
			// Place it at the kth position
			nums[k] = nums[i]
			// Move the pointer forward
			k++
		}
	}

	// k represents the number of unique elements
	return k
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicatesMaxTwoChatGPT(nums []int) int {
	// Edge case: empty array
	if len(nums) <= 2 {
		return len(nums)
	}

	// Initialize a pointer for placing unique elements
	k := 2

	// Iterate through the input array starting from the third element
	for i := 2; i < len(nums); i++ {
		// If the current element is different from the (k-2)th element
		if nums[i] != nums[k-2] {
			// Place it at the kth position
			nums[k] = nums[i]
			// Move the pointer forward
			k++
		}
	}

	// k represents the number of unique elements (at most two occurrences)
	return k
}
