package turing

func countPairs(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	pairsCount := 0
	for _, count := range counts {
		if count%2 != 0 {
			return -1
		}
		pairsCount += count / 2
	}

	return pairsCount
}
