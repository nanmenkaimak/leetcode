package main

func missingNumber(nums []int) int {
	res := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
	}

	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			return i
		}
	}
	return 0
}
