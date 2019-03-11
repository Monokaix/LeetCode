package main

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(height(root.Left)-height(root.Right))) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

}

// compute tree height
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h1 := height(root.Left)
	h2 := height(root.Right)
	if h1 < h2 {
		return h2 + 1
	}
	return h1 + 1
}
