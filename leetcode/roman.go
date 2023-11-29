package leetcode

import (
	"strings"
)

func RomanToInt(s string) int {
	lookup := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	strSlice := strings.Split(s, "")
	sum := 0

	for i := 0; i < len(strSlice); i++ {
		currentValue := lookup[strSlice[i]]

		if i == 0 {
			sum += currentValue
			continue
		}

		prevValue := lookup[strSlice[i-1]]

		if currentValue > prevValue {
			sum -= prevValue * 2
		}
		sum += currentValue

	}

	return sum
}
