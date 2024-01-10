package search

func Linear(arr []int, needle int) bool {
	for i := 0; i < len(arr); i++ {
		if needle == arr[i] {
			return true
		}
	}
	return false
}
