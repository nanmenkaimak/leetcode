package main

func maxProfit(prices []int) int {
	profit := 0
	minNumber := prices[0]
	for i := 1; i < len(prices); i++ {
		profitDay := prices[i] - minNumber
		if profitDay > profit {
			profit = profitDay
		}
		if minNumber > prices[i] {
			minNumber = prices[i]
		}
	}
	return profit
}
