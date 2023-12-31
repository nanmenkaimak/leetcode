package main

import "math"

func maxSubArray(nums []int) int {
	curMax, res := 0, math.MinInt

	for i := 0; i < len(nums); i++ {
		curMax = maxn(curMax+nums[i], nums[i])
		res = maxn(res, curMax)
	}
	return res
}

func maxn(nums1, nums2 int) int {
	if nums1 > nums2 {
		return nums1
	}
	return nums2
}
