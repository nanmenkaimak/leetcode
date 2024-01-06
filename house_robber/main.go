package main

func rob(nums []int) int {
	prev1, prev2 := 0, 0

	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev1, prev2+num)
		prev2 = tmp
	}
	return prev1
}
