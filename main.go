package main

import (
	"fmt"
	"math"

	"github.com/padunk/go-dsa/leetcode/window"
	// "sort"
)

func main() {
	// people := []Person{
	// 	{"Eve", 18},
	// 	{"Alice", 30},
	// 	{"Bob", 25},
	// }

	// sort.Sort(ByAge(people))
	// fmt.Println(people)

	sample := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	complete := window.GPTMinSubArrayLen(213, sample)
	fmt.Println(complete)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func selectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		smallestIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}

		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
	}
}

func quickSort(arr []int) []int {
	// leave
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) - 1
	pivot := arr[pivotIndex]
	arr = append(arr[:pivotIndex], arr[pivotIndex+1:]...)

	var left, right []int

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)

}

// ! still don't get this
func twoCrystalProblems(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1
}

// ! this search functions works in sorted slice
func linearSearch(arr []int, n int) int {
	result := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			result = i
			break
		}
	}
	return result
}

func binarySearch(haystack []int, n int) bool {
	lo, hi := 0, len(haystack)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := haystack[m]

		if v == n {
			return true
		} else if v > n {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
