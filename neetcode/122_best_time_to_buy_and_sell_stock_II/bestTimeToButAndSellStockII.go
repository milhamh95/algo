package best_time_to_buy_and_sell_stock_II

// O(n ^ 2)
// totalProfit = 0

func maxProfit(prices []int) int {
	totalProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit := prices[i] - prices[i-1]
			totalProfit = totalProfit + profit
		}
	}

	return totalProfit
}
