package search

func findPeakElement(nums []int) int {
	low, hi := 0, len(nums)-1

	for low < hi {
		mid := low + (hi-low)/2

		// If mid is greater than its right neighbor, a peak might exist to the left of mid
		if nums[mid] > nums[mid+1] {
			hi = mid
		} else {
			// Otherwise, a peak might exist to the right of mid
			low = mid + 1
		}
	}

	// Left will eventually converge to the peak element
	return low
}
