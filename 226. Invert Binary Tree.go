package main

//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil{
//		return nil
//	}
//	root.Left,root.Right= root.Right,root.Left
//	invertTree(root.Left)
//	invertTree(root.Right)
//	return root
//}

// 方法二
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 分别反转左右子树
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	// 然后左右互换两个子树
	root.Left = right
	root.Right = left

	return root
}
