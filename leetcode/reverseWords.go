package leetcode

import (
	"slices"
	"strings"
)

func ReverseWords(s string) string {
	s = strings.Trim(s, " ")
	strs := strings.Split(s, " ")
	result := []string{}
	for i := 0; i < len(strs); i++ {
		if strs[i] != "" {
			result = append(result, strings.Trim(strs[i], " "))
		}
	}
	slices.Reverse(result)
	return strings.Join(result, " ")
}

func OtherReverseWords(s string) string {
	strs := strings.Split(s, " ")
	result := []string{}
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			result = append(result, strings.Trim(strs[i], " "))
		}
	}
	return strings.Join(result, " ")
}
