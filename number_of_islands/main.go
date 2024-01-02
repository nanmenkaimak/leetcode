package main

func numIslands(grid [][]byte) int {
	cnt := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt++
				kk(grid, i, j)
			}
		}
	}
	return cnt
}

func kk(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return
	}

	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'

	kk(grid, i+1, j)
	kk(grid, i-1, j)
	kk(grid, i, j+1)
	kk(grid, i, j-1)
}
