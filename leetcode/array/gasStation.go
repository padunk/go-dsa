package leetcode

// https://leetcode.com/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
func CanCompleteCircuit(gas []int, cost []int) int {
	gasNeeded := make([]int, len(gas))

	for i := 0; i < len(gas); i++ {
		gasNeeded[i] = gas[i] - cost[i]
	}

	return -1
}

func canCompleteCircuitGPT(gas []int, cost []int) int {
	totalGas := 0
	currentGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		currentGas += diff
		totalGas += diff

		if currentGas < 0 {
			currentGas = 0
			start = i + 1
		}
	}

	if totalGas < 0 {
		return -1
	}

	return start % len(gas)
}
