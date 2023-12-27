package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	} else if kkk(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func kkk(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	} else if root.Val == subRoot.Val {
		return kkk(root.Right, subRoot.Right) && kkk(root.Left, subRoot.Left)
	}
	return false
}
