package main

func isSymmetric(root *TreeNode) bool {
	return isMirror(root,root)
}
func isMirror(p *TreeNode, q *TreeNode) bool  {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
	} else {
		return false
	}
}