package besttimetobuyandsellstock

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	idx1 := 0
	idx2 := 1
	maxProfit := 0

	for idx2 < len(prices) {
		if prices[idx1] < prices[idx2] {
			profit := prices[idx2] - prices[idx1]

			if maxProfit < profit {
				maxProfit = profit
			}
		} else {
			idx1 = idx2
		}

		idx2++
	}

	return maxProfit
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := math.MinInt

	for _, v := range prices {
		money := v - min
		if money > max {
			max = money
		}

		if v < min {
			min = v
		}
	}

	return max
}

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	//for _, v := range prices {
	//	profit := v - minPrice
	//	maxProfit = max(maxProfit, profit)
	//	minPrice = min(minPrice, v)
	//}

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		maxProfit = max(maxProfit, profit)
		minPrice = min(minPrice, prices[i])
	}

	return maxProfit
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
