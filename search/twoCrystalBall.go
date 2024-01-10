package search

import "math"

func TwoCyrstalBall(breaks []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jump
	for ; i < len(breaks); i += jump {
		if breaks[i] {
			break
		}
	}

	j := i - jump
	for ; j < i; j++ {
		if breaks[j] {
			return j
		}
	}

	return -1

}

// imagined two balls drop from high above building
// determined the best way to find both floor where it break
func twoCrystalProblems(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1
}
