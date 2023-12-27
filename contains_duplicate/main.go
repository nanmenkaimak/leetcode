package main

func containsDuplicate(nums []int) bool {
	duplicate := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		duplicate[nums[i]]++
		if duplicate[nums[i]] > 1 {
			return true
		}
	}
	return false
}