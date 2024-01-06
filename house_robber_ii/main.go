package main

import "slices"

func rob(nums []int) int {
	if len(nums) <= 3 {
		return slices.Max(nums)
	}

	return max(sum(0, len(nums)-1, nums), sum(1, len(nums), nums))
}

func sum(start int, end int, nums []int) int {
	prev_prev := nums[start]
	prev := max(nums[start], nums[start+1])

	for i := start + 2; i < end; i++ {
		prev_prev, prev = prev, max(prev_prev+nums[i], prev)
	}

	return prev
}
