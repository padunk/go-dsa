package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
func maxProfit(prices []int) int {
	minPrice, profit := prices[0], 0

	for i, price := range prices {
		if price < minPrice {
			minPrice = price
			if i == len(prices)-1 && profit == 0 {
				return 0
			}
		}
		if price > minPrice && price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}

func maxProfitChatGPT(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func maxProfitII(prices []int) int {
	profit := 0
	buyPrice := -1

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] && buyPrice == -1 {
			buyPrice = prices[i]
		}
		if prices[i] > prices[i+1] && buyPrice > -1 {
			profit += prices[i] - buyPrice
			buyPrice = -1
		}
	}

	if buyPrice > -1 {
		profit += prices[len(prices)-1] - buyPrice
	}

	return profit
}

func maxProfitIIChatGPT(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
