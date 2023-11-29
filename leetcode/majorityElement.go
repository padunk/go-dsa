package leetcode

func majorityElement(nums []int) int {
	m, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			m = num
			count = 1
		} else if m == num {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if m == num {
			count++
		}
	}

	if count > len(nums)/2 {
		return m
	} else {
		return -1
	}
}
