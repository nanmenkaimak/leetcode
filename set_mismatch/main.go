package set_mismatch

func findErrorNums(nums []int) []int {
	var res []int

	maps := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		maps[nums[i]]++
		if maps[nums[i]] > 1 {
			res = append(res, nums[i])
		}
	}

	for i := 1; i <= len(nums); i++ {
		if maps[i] == 0 {
			res = append(res, i)
			return res
		}
	}
	return []int{}
}
