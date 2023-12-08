package leetcode

import (
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

	return strs
}

// GPT Full Justify
func GPTFullJustify(words []string, maxWidth int) []string {
	result := []string{}
	startIndex := 0

	for startIndex < len(words) {
		endIndex := findEndIndex(words, startIndex, maxWidth)

		if endIndex == startIndex || endIndex == len(words)-1 {
			// If only one word can fit on a line, left-justify it.
			result = append(result, leftJustify(words[startIndex:endIndex+1], maxWidth))
		} else {
			// Otherwise, fully justify the line.
			result = append(result, fullJustifyLine(words[startIndex:endIndex+1], maxWidth))
		}

		startIndex = endIndex + 1
	}

	return result
}

func findEndIndex(words []string, startIndex, maxWidth int) int {
	currentLength := len(words[startIndex])
	endIndex := startIndex + 1

	for endIndex < len(words) && currentLength+len(words[endIndex])+1 <= maxWidth {
		currentLength += len(words[endIndex]) + 1
		endIndex++
	}

	return endIndex - 1
}

func leftJustify(words []string, maxWidth int) string {
	return strings.Join(words, " ") + strings.Repeat(" ", maxWidth-len(strings.Join(words, " ")))
}

func fullJustifyLine(words []string, maxWidth int) string {
	wordCount := len(words)
	totalSpaces := maxWidth - len(strings.Join(words, ""))
	extraSpaces := totalSpaces % (wordCount - 1)
	spacePerWord := totalSpaces / (wordCount - 1)

	line := words[0]

	for i := 1; i < wordCount; i++ {
		spaces := spacePerWord
		if i <= extraSpaces {
			spaces++
		}
		line += strings.Repeat(" ", spaces) + words[i]
	}

	return line
}

// other solutions
func OtherFullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	line := []string{}
	someWordWidth := 0

	appendLeftJustified := func() {
		strLine := line[0]
		for _, word := range line[1:] {
			strLine += " " + word
		}
		strLine += strings.Repeat(" ", maxWidth-len(strLine))
		lines = append(lines, strLine)
	}

	appendStringLine := func() {
		if len(line) == 1 {
			appendLeftJustified()
			return
		}
		strLine := line[0]
		spaceSize := (maxWidth - someWordWidth) / (len(line) - 1)
		remain := (maxWidth - someWordWidth) % (len(line) - 1)
		space := strings.Repeat(" ", spaceSize+1)
		for index, word := range line[1:] {
			if index == remain {
				space = space[1:]
			}
			strLine += space
			strLine += word
		}
		lines = append(lines, strLine)
	}

	for _, word := range words {
		if someWordWidth+len(line)+len(word) > maxWidth {
			appendStringLine()
			line = []string{}
			someWordWidth = 0
		}
		line = append(line, word)
		someWordWidth += len(word)
	}
	appendLeftJustified()
	return lines
}
