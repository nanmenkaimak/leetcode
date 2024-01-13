package main

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	res, cur := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			cur++
		}
		for cur > k {
			if nums[left] == 0 {
				cur--
			}
			left++
		}
		right++
		res = max(res, right-left)
	}
	return res
}
