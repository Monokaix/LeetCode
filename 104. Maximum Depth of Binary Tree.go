package main

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left < right{
		return right+1
	}
	return left+1
}
