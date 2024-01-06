package main

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	sum(0, make([]int, 0), 0, target, &res, candidates)
	return res
}

func sum(i int, cur []int, total int, target int, res *[][]int, candidates []int) {
	if total == target {
		cpyTmp := make([]int, len(cur))
		copy(cpyTmp, cur)
		*res = append(*res, cpyTmp)
	}
	if i >= len(candidates) || total > target {
		return
	}
	for idx := i; idx < len(candidates); idx++ {
		cur = append(cur, candidates[idx])
		sum(idx, cur, total+candidates[idx], target, res, candidates)
		cur = cur[:len(cur)-1]
	}
}
