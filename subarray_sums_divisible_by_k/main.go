package main

func subarraysDivByK(nums []int, k int) int {
	res := 0
	prefixSum := 0
	maps := make(map[int]int)
	maps[0] = 1

	for i := range nums {
		prefixSum += nums[i]
		prefixSum = ((prefixSum % k) + k) % k

		res += maps[prefixSum]
		maps[prefixSum]++
	}
	return res
}
