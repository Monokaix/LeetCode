package main

import "math"

func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt32
	pre := -1
	inorder(root, &pre, &res)
	return res
}

func inorder(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}

	inorder(root.Left, pre, min)
	if *pre != -1 {
		*min = Min(*min, root.Val-*pre)
	}
	*pre = root.Val
	inorder(root.Right, pre, min)
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
