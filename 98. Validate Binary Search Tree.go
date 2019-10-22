package main

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs_left(root.Left, root.Val) && dfs_right(root.Right, root.Val) && isValidBST(root.Left) && isValidBST(root.Right)

}
func dfs_left(root *TreeNode, root_val int) bool {
	if root == nil {
		return true
	}
	if root.Val >= root_val {
		return false
	}
	return dfs_left(root.Left, root_val) && dfs_left(root.Right, root_val)
}

func dfs_right(root *TreeNode, root_val int) bool {
	if root == nil {
		return true
	}
	if root.Val <= root_val {
		return false
	}
	return dfs_right(root.Left, root_val) && dfs_right(root.Right, root_val)
}

//方法二 cmp是递增的,因为搜索树要满足中序序列递增
func isValidBST2(root *TreeNode) bool {
	cmp := math.MinInt64
	return isValidBST3(root, &cmp)
}
func isValidBST3(root *TreeNode, cmp *int) bool {
	if root == nil {
		return true
	}
	if isValidBST3(root.Left, cmp) {
		if *cmp < root.Val {
			*cmp = root.Val
			return isValidBST3(root.Right, cmp)
		}
	}
	return false
}
