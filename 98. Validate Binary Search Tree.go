package main

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
