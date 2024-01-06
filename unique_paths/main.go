package main

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	for i := range row {
		row[i] = 1
	}

	for i := 0; i < m-1; i++ {
		newRow := make([]int, n)
		for k := range newRow {
			newRow[k] = 1
		}
		for j := n - 2; j >= 0; j-- {
			newRow[j] = newRow[j+1] + row[j]
		}
		row = newRow
	}
	return row[0]
}
