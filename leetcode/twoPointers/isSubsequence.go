package twopointers

func IsSubsequence(s string, t string) bool {
	sLength := len(s)
	tLength := len(t)

	if sLength == 0 {
		return true
	}
	if sLength > tLength {
		return false
	}
	i, j := 0, 0

	for ; i < sLength && j < tLength; j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == sLength
}

func OtherIsSubsequence(s string, t string) bool {
	i, sIdx := 0, 0
	for ; i < len(t) && sIdx < len(s); i++ {
		if s[sIdx] == t[i] {
			sIdx++
		}
	}
	return sIdx == len(s)

}
