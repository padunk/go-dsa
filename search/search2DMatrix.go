package search

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
	upperHigh := len(matrix)
	upperLow := 0

	for upperLow < upperHigh {
		upperMid := upperLow + (upperHigh-upperLow)/2
		subMatrix := matrix[upperMid]
		high := len(subMatrix)
		lo := 0

		if target < subMatrix[0] {
			upperHigh = upperMid
			continue
		}
		if target > subMatrix[high-1] {
			upperLow = upperMid + 1
			continue
		}

		// search subMatrix
		for lo < high {
			mid := lo + (high-lo)/2
			if subMatrix[mid] == target {
				return true
			}
			if target < subMatrix[mid] {
				high = mid
			} else {
				lo = mid + 1
			}
		}
		fmt.Println(lo, high)
		if lo >= high {
			break
		}
	}

	return false
}

func searchMatrixOthers(matrix [][]int, target int) bool {
	hi, low := len(matrix)*len(matrix[0])-1, 0

	for low <= hi {
		mid := (low + hi) / 2
		x := mid / len(matrix[0])
		y := mid % len(matrix[0])
		if matrix[x][y] > target {
			hi = mid - 1
		} else if matrix[x][y] < target {
			low = mid + 1
		} else {
			return true
		}
	}

	return false
}
