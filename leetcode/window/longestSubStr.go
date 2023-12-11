package window

import (
	"slices"
)

func LengthOfLongestSubstring(s string) int {
	temp := []rune{}
	maxLength := 0

	for _, v := range s {
		if slices.Contains(temp, v) {
			idx := slices.Index(temp, v)
			temp = slices.Delete(temp, 0, idx+1)
		}
		temp = append(temp, v)
		maxLength = max(maxLength, len(temp))
	}

	return maxLength
}

func otherLength(s string) int {
	charIndexMap := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if index, ok := charIndexMap[s[right]]; ok && index >= left {

			left = index + 1
		}
		charIndexMap[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
