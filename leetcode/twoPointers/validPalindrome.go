package twopointers

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	result := reg.ReplaceAllString(s, "")
	result = strings.ToLower(result)
	j := len(result) - 1

	for i := 0; i < len(result) && j >= 0; i++ {
		if result[i] != result[j] {
			return false
		}
		j--
	}

	return true
}
