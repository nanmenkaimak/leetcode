package minimum_absolute_difference

import (
	"math"
	"slices"
)

func minimumAbsDifference(arr []int) [][]int {
	var res [][]int
	minNum := math.MaxInt

	slices.Sort(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			minNum = min(minNum, arr[i]-arr[i-1])
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minNum {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
