package main

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, node := getDepthAndNode(root)
	return node
}

func getDepthAndNode(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	// 递归寻找左子树的最大深度及其对应的公共祖先
	leftDepth, leftNode := getDepthAndNode(root.Left)
	rightDepth, rightNode := getDepthAndNode(root.Right)

	// 相等时直接返回，即当前节点即为要找的公共祖先
	if leftDepth == rightDepth {
		return leftDepth + 1, root
	}
	if leftDepth > rightDepth {
		return leftDepth + 1, leftNode
	} else {
		return rightDepth + 1, rightNode
	}
}
