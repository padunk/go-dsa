package search

import (
	"math"
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortedArray := append([]int{}, nums1...)
	sortedArray = append(sortedArray, nums2...)
	slices.Sort(sortedArray)

	length := len(sortedArray)
	mid := length / 2
	if length%2 == 0 {
		return (float64(sortedArray[mid-1]) + float64(sortedArray[mid])) / 2
	} else {
		return float64(sortedArray[mid])
	}
}

func findMedianSortedArraysGPT(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX
		maxLeftX := math.MinInt64
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}
		minRightX := math.MaxInt64
		if partitionX != x {
			minRightX = nums1[partitionX]
		}
		maxLeftY := math.MinInt64
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}
		minRightY := math.MaxInt64
		if partitionY != y {
			minRightY = nums2[partitionY]
		}
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return -1
}
