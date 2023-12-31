package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix_sum := 1
	postfix_sum := 1

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = prefix_sum
		prefix_sum *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		res[i] *= postfix_sum
		postfix_sum *= nums[i]
	}
	return res
}
