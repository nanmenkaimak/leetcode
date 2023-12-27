package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, prev2, res := 1, 2, 0

	for i := 3; i <= n; i++ {
		res = prev + prev2
		prev = prev2
		prev2 = res
	}
	return res
}
