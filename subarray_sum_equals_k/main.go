package main

func subarraySum(nums []int, k int) int {
	prefixSum := make([]int, len(nums))
	count := map[int]int{0: 1}

	res := 0
	previousPrefix := 0
	for i := 0; i < len(nums); i++ {
		prefixSum[i] = nums[i] + previousPrefix
		previousPrefix = prefixSum[i]

		remaining := prefixSum[i] - k
		res += count[remaining]

		count[prefixSum[i]]++
	}

	return res
}
