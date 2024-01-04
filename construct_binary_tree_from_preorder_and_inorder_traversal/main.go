package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	val := preorder[0]
	ind := 0

	for ind < len(inorder) && inorder[ind] != val {
		ind++
	}

	ans := new(TreeNode)
	ans.Val = val

	ans.Left = buildTree(preorder[1:], inorder[:ind])
	ans.Right = buildTree(preorder[ind+1:], inorder[ind+1:])

	return ans
}
