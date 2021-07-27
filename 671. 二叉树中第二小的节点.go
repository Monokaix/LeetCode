package main

func findSecondMinimumValue(root *TreeNode) int {
	return find(root, root.Val)
}

func find(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}

	// 此时root.Val就是要求的值，因为另外一个节点的值和根节点相同
	if root.Val > val {
		return root.Val
	}
	l := find(root.Left, val)
	r := find(root.Right, val)
	if l == -1 {
		return r
	}
	if r == -1 {
		return l
	}
	return min(l, r)
}
