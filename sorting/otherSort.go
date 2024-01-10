package sorting

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
