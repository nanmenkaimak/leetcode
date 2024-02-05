package sum_of_root_to_leaf_binary_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return kk(root, 0)
}

func kk(root *TreeNode, ans int) int {
	if root == nil {
		return 0
	}

	ans = ans*2 + root.Val

	if root.Left == nil && root.Right == nil {
		return ans
	}

	return kk(root.Left, ans) + kk(root.Right, ans)
}
