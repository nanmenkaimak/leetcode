package main

import "math"

func maxProduct(nums []int) int {
	res := math.MinInt

	for i := 0; i < len(nums); i++ {
		res = max(res, nums[i])
	}

	curMax, curMin := 1, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			curMax, curMin = 1, 1
			continue
		}

		maxNum := curMax * nums[i]

		curMax = max(max(maxNum, curMin*nums[i]), nums[i])
		curMin = min(min(maxNum, curMin*nums[i]), nums[i])

		res = max(max(res, curMax), curMin)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
