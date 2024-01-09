package main

func canJump(nums []int) bool {
	num := 0

	for i := 0; i < len(nums); i++ {
		if i > num {
			return false
		}
		num = max(num, i+nums[i])
	}
	return true
}
