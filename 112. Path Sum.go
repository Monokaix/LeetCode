package main

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}
