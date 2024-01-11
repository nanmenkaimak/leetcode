package main

import "math"

func numSquares(n int) int {
	res := make([]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = math.MaxInt
		for j := 1; j <= i; j++ {
			pow := j * j
			if i-pow < 0 {
				break
			}

			res[i] = min(res[i], res[i-pow]+1)
		}
	}
	return res[n]
}
