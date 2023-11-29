package leetcode

import (
	"fmt"
	"strings"
)

func FullJustify(words []string, maxWidth int) []string {
	concateWord := words[0]
	wordLength := len(concateWord)

	for i := 1; i < len(words); i++ {
		if wordLength+len(words[i])+1 > maxWidth {
			concateWord += "\n" + words[i]
			wordLength = len(words[i])
		} else {
			concateWord += " " + words[i]
			wordLength += len(words[i]) + 1
		}
	}

	strs := strings.Split(concateWord, "\n")
	for i := 0; i < len(strs)-1; i++ {
		fmt.Println(len(strs[i]))
	}

	return []string{""}
}
