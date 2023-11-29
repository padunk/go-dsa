package leetcode

func canJump(nums []int) bool {
	length := len(nums)
	maxReach := 0

	for i := 0; i < length; {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if maxReach >= length-1 {
			return true
		}
		i++
	}

	return false
}

// https://leetcode.com/problems/jump-game-ii/submissions/1101189537/?envType=study-plan-v2&envId=top-interview-150
func jump(nums []int) int {
	length := len(nums)
	jumps := 0
	farthest := 0
	currentEnd := 0

	for i := 0; i < length-1; i++ {
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}
