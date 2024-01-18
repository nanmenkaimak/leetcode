package main

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	val := 1

	for left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = val
			val++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[i][right] = val
			val++
		}
		right--

		for i := right; i >= left; i-- {
			res[bottom][i] = val
			val++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res[i][left] = val
			val++
		}
		left++
	}
	return res
}
