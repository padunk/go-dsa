package turing

import (
	"strconv"
	"strings"
)

func perfectInteger(x, y int) []int {
	res := []int{}

	for i := x; i <= y; i++ {
		v := strings.Split(strconv.Itoa(i), "")
		check := make([]bool, len(v))

		for j := 0; j < len(v); j++ {
			vInt, _ := strconv.Atoi(v[j])
			if vInt == 0 {
				break
			}
			if i%vInt == 0 {
				check[j] = true
			}
		}

		checkRes := true
		for k := 0; k < len(check); k++ {
			checkRes = checkRes && check[k]
		}
		if checkRes {
			res = append(res, i)
		}
	}

	return res
}

// ChatGPT
func PerfectInteger(x, y int) []int {
	var res []int

	for i := x; i <= y; i++ {
		if isDivisibleByDigits(i) {
			res = append(res, i)
		}
	}

	return res
}

func isDivisibleByDigits(num int) bool {
	digits := getDigits(num)
	for _, digit := range digits {
		if digit == 0 || num%digit != 0 {
			return false
		}
	}
	return true
}

func getDigits(num int) []int {
	var digits []int
	strNum := strconv.Itoa(num)

	for _, char := range strNum {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}

	return digits
}
