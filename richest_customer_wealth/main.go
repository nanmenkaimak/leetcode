package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	maxNum := 0

	for i := 0; i < len(accounts); i++ {
		res := 0
		for j := 0; j < len(accounts[0]); j++ {
			res += accounts[i][j]
		}
		maxNum = max(maxNum, res)
	}
	return maxNum
}
