package main

func searchRange(nums []int, target int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
			break
		}
	}

	if len(res) == 0 {
		return []int{-1, -1}
	}

	for i := len(nums) - 1; i >= res[0]; i-- {
		if nums[i] == target {
			res = append(res, i)
			break
		}
	}

	return res
}
