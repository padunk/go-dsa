package leetcode

func Trap(heights []int) int {
	n := len(heights)
	maxHeight := 0
	leftToRight := make([]int, n)
	rightToLeft := make([]int, n)

	leftToRight[0] = 0
	for i := 1; i < n; i++ {
		if maxHeight < max(heights[i], heights[i-1]) {
			maxHeight = max(heights[i], heights[i-1])
		}
		if heights[i] < maxHeight {
			leftToRight[i] = maxHeight - heights[i]
		}
	}

	rightToLeft[n-1] = 0
	maxHeight = 0
	for i := n - 2; i >= 0; i-- {
		if maxHeight < max(heights[i], heights[i+1]) {
			maxHeight = max(heights[i], heights[i+1])
		}
		if heights[i] < maxHeight {
			rightToLeft[i] = maxHeight - heights[i]
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += min(leftToRight[i], rightToLeft[i])
	}
	return sum

}

func TrapGPT(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Calculate maximum height from the left for each position
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// Calculate maximum height from the right for each position
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// Calculate trapped water at each position based on the minimum of leftMax and rightMax
	trappedWater := 0
	for i := 1; i < n-1; i++ {
		minHeight := min(leftMax[i], rightMax[i])
		trappedWater += max(0, minHeight-height[i])
	}

	return trappedWater
}
