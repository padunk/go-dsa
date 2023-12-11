package window

import (
	"strings"
)

func FindSubstring(s string, words []string) []int {
	output := []int{}
	end := len(words) * len(words[0])

	for i := 0; i < len(s); i++ {
		if i+end > len(s) {
			break
		}

		checkStr := s[i : i+end]
		isValid := true
		for j := 0; j < len(words); j++ {
			idx := strings.Index(checkStr, words[j])
			if idx != -1 {
				checkStr = strings.Replace(checkStr, words[j], "", 1)
			} else {
				isValid = false
				break
			}
		}
		if isValid {
			output = append(output, i)
		}
		isValid = true
	}

	return output
}

// CHATGPT
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount
	result := []int{}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i <= len(s)-totalLen; i++ {
		substr := s[i : i+totalLen]
		subMap := make(map[string]int)

		for j := 0; j < totalLen; j += wordLen {
			word := substr[j : j+wordLen]
			subMap[word]++
		}

		if mapsAreEqual(wordMap, subMap) {
			result = append(result, i)
		}
	}

	return result
}

func mapsAreEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, count1 := range map1 {
		if count2, found := map2[key]; !found || count1 != count2 {
			return false
		}
	}

	return true
}
