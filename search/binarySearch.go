package search

func Binary(arr []int, needle int) bool {
	lo := 0
	hi := len(arr)

	for lo < hi {
		mid := lo + (hi-lo)/2
		value := arr[mid]
		if needle == value {
			return true
		} else if value > needle {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return false
}
