package leetcode

func strStr(haystack, needle string) int {

	if haystack == "" {
		return -1
	}
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
