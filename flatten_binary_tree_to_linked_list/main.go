package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var nums []int
	kk(root, &nums)

	f(root, nums, 0)
}

func kk(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	kk(root.Left, nums)
	kk(root.Right, nums)
}

func f(root *TreeNode, nums []int, i int) int {
	if i >= len(nums) {
		return i
	}
	root.Left = nil
	root.Val = nums[i]
	i++
	if i < len(nums) {
		root.Right = &TreeNode{}
		i = f(root.Right, nums, i)
	}
	return i
}
