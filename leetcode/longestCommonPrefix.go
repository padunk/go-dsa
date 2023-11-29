package leetcode

import (
	"fmt"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Take the first string as the initial common prefix
	prefix := strs[0]

	// Iterate through the remaining strings
	for _, s := range strs[1:] {
		// Adjust the prefix based on the current string
		for !strings.HasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
			fmt.Println(prefix)
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

// other solution:
func otherSolution(strs []string) string {
	prefix := ""

	for _, v := range strs[0] {
		add := true
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], prefix+string(v)) {
				add = false
			}
		}

		if add {
			prefix += string(v)
		} else {
			break
		}
	}

	return prefix
}
