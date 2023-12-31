package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	prev, cnt := 0, 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[prev][1] {
			prev = i
			cnt++
		}
	}
	return len(intervals) - cnt
}
