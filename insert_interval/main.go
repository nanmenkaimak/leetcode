package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[1] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			return ans
		} else if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
		} else {
			newInterval = []int{
				min(newInterval[0], intervals[i][0]),
				max(newInterval[1], intervals[i][1]),
			}
		}
	}

	ans = append(ans, newInterval)

	return ans
}
