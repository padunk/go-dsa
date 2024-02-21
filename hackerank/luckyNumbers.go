package hackerank

import (
	"math"
)

func LuckyNumber(a, b int) []int {
	var result []int

	for i := a; i <= b; i++ {
		if isPrime(sumOfDigits(i, 1.0)) && isPrime(sumOfDigits(i, 2.0)) {
			result = append(result, i)
		}
	}

	return result
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	sqrt := int(math.Floor(math.Sqrt(float64(num))))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// ChatGPT
func sumOfDigits(n int, pow float64) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), pow))
		n /= 10
	}

	return sum
}
