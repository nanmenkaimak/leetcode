package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1

		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right -= 1
			} else if sum < 0 {
				left += 1
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return res
}
