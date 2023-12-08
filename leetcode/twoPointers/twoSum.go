package twopointers

import "fmt"

func TwoSum(numbers []int, target int) []int {
	length := len(numbers)
	result := make([]int, 2)

	i, j := 0, length-1
	for i < length && j >= 0 {
		fmt.Println(i, j)
		lookFor := target - numbers[i]
		if numbers[j] == lookFor {
			result[0] = i + 1
			result[1] = j + 1
			break
		}
		if numbers[j] < lookFor {
			i++
		} else {
			j--
		}
	}

	return result
}

// other solutions
func OtherTwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
