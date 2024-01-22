package path_sum_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	maps := make(map[int]int)
	maps[0] = 1
	return kk(root, targetSum, 0, maps)
}

func kk(root *TreeNode, targetSum int, prefixSum int, maps map[int]int) int {
	if root == nil {
		return 0
	}
	prefixSum += root.Val
	ans := maps[prefixSum-targetSum]
	maps[prefixSum]++

	ans += kk(root.Left, targetSum, prefixSum, maps)
	ans += kk(root.Right, targetSum, prefixSum, maps)

	maps[prefixSum]--
	return ans
}
