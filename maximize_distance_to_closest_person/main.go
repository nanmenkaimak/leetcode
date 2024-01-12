package main

func maxDistToClosest(seats []int) int {
	res, left, idx := 0, -1, 0

	for idx < len(seats) {
		if left == -1 && seats[idx] == 1 {
			res = max(res, idx)
			left = idx
		} else if idx == len(seats)-1 && seats[idx] == 0 {
			res = max(res, idx-left)
		} else if seats[idx] == 1 {
			res = max(res, (idx-left)/2)
			left = idx
		}
		idx++
	}
	return res
}
