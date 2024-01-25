package arranging_coins

func arrangeCoins(n int) int {
	res, i := 0, 1

	if n <= 1 {
		return n
	}

	for n > 0 {
		res++
		n -= i
		i++
	}
	return res
}
