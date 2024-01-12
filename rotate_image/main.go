package main

import "slices"

func rotate(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		slices.Reverse(matrix[i])
	}
}
