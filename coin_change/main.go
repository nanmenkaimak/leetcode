package main

func coinChange(coins []int, amount int) int {
	moves := make([]int, amount+1)
	for i := range moves {
		moves[i] = amount + 1
	}
	moves[0] = 0

	for i := 1; i < len(moves); i++ {
		for c := 0; c < len(coins); c++ {
			if i-coins[c] >= 0 {
				moves[i] = min(moves[i], moves[i-coins[c]]+1)
			}
		}
	}
	if moves[amount] != amount+1 {
		return moves[amount]
	}
	return -1
}
