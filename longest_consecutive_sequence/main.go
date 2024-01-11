package main

import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)

	ans := 0
	curAns := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			curAns++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			ans = max(ans, curAns+1)
			curAns = 0
		}
	}
	return max(ans, curAns+1)
}
