package leetcode

import "strings"

func LengthOfLastWord(s string) int {
	str := strings.Split(strings.Trim(s, " "), " ")
	lastWord := str[len(str)-1]
	return len(lastWord)
}
