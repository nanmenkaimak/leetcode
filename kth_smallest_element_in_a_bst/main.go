package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func fun(root *TreeNode, val *[]int) {
	if root == nil {
		return
	}
	fun(root.Left, val)
	*val = append(*val, root.Val)
	fun(root.Right, val)
}

func kthSmallest(root *TreeNode, k int) int {
	var val []int

	fun(root, &val)
	return val[k-1]
}
