package two_sum_IV

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func dfs(root *TreeNode, k int, maps map[int]bool) bool {
	if root == nil {
		return false
	}

	if maps[k-root.Val] {
		return true
	}
	maps[root.Val] = true

	return dfs(root.Left, k, maps) || dfs(root.Right, k, maps)
}
