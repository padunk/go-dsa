package sorting

func myQuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	doQuickSort(&arr, 0, len(arr)-1)
}

func doQuickSort(arr *[]int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIndex := partition(arr, lo, hi)
	doQuickSort(arr, lo, pivotIndex-1)
	doQuickSort(arr, pivotIndex+1, hi)
}

func partition(arr *[]int, lo, hi int) int {
	pivot := (*arr)[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if (*arr)[i] <= pivot {
			idx++
			(*arr)[i], (*arr)[idx] = (*arr)[idx], (*arr)[i]
		}
	}

	idx++
	if (*arr)[idx] > (*arr)[hi] {
		(*arr)[idx], (*arr)[hi] = (*arr)[hi], (*arr)[idx]
	}
	return idx
}
