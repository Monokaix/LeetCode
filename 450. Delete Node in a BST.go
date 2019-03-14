package main

func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode2(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode2(root.Right, key)
		return root
	} else { //找到删除节点
		if root.Left == nil { //待删除节点左子树为空，之间返回右子树作为新的根
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			successor := findLeft(root.Right)
			successor.Right = deleteLeft(root.Right)
			successor.Left = root.Left
			return successor
		}
	}
}

// 找到右子树最左边节点
func findLeft(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return findLeft(node.Left)
}

//将最左节点置空并返回最左节点父节点作为
func deleteLeft(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteLeft(node.Left)
	return node
}
