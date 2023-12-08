package leetcode

import "sort"

// https://leetcode.com/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150

// sort the citations
func hIndexSort(citations []int) int {
	count := 0
	// sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := 0; i < len(citations); i++ {
		if count < citations[i] {
			count++
		}
	}

	return count
}

// without sort the citations
func hIndex(citations []int) int {
	n := len(citations)

	// Create a histogram of citation counts
	histogram := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			histogram[n]++
		} else {
			histogram[citation]++
		}
	}

	// Iterate through the histogram to find the H-index
	papers := 0
	for i := n; i >= 0; i-- {
		papers += histogram[i]
		if papers >= i {
			return i
		}
	}

	return 0
}
