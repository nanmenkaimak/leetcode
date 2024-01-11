package main

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	l, r := 0, len(nums)-1
	i := len(nums) - 1

	for l <= r {
		if absInt(nums[l]) > absInt(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
		i--
	}
	return res
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
