package main

func findMin(nums []int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if minNum > nums[i] {
			return nums[i]
		}
	}
	return minNum
}
