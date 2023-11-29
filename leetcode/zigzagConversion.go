package leetcode

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	if numRows < 2 || numRows >= len(s) {
		return s
	}

	strs := strings.Split(s, "")
	strLength := len(strs)
	result := ""
	jump := 2*numRows - 2
	vertJump := jump
	hortJump := jump

	for i := 0; i < numRows; i++ {
		j := i
		k := 0
		hortJump = vertJump

		for j < strLength {
			result += strs[j]
			if k > 0 && hortJump < jump {
				hortJump = jump - hortJump
			}

			j += hortJump
			k++
		}
		vertJump -= 2

		if vertJump == 0 {
			vertJump = jump
		}
	}

	return result
}

// other solutions
func OtherConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]byte, numRows)

	row, step := 0, 1
	for _, v := range s {
		rows[row] = append(rows[row], byte(v))

		if row+step == numRows || row+step == -1 {
			step = -step
		}

		row += step
	}

	res := make([]byte, 0, len(s))
	for _, row := range rows {
		res = append(res, row...)
	}

	return string(res)
}

// chatGPT
// this using step, where step use to control the index
func GPTConvert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]string, numRows)
	index, step := 0, 1

	for i := range s {
		result[index] += string(s[i])

		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}

		index += step
	}
	fmt.Println(result)

	return strings.Join(result, "")
}
