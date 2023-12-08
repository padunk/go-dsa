package twopointers

func MaxArea(height []int) int {
	area := 0
	i, j := 0, len(height)-1

	for i <= j {
		area = max(area, (j-i)*min(height[j], height[i]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return area
}
