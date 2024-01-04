package main

func lengthOfLIS(nums []int) int {
	moves := make([]int, len(nums))
	for i := range moves {
		moves[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				moves[i] = max(moves[i], 1+moves[j])
			}
		}
	}

	res := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] > res {
			res = moves[i]
		}
	}

	return res
}
