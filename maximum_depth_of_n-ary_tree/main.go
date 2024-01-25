package maximum_depth_of_n_ary_tree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	res := 0
	for _, v := range root.Children {
		res = max(res, maxDepth(v))
	}
	return res + 1
}
