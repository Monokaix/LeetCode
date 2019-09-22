package main

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(inorder) == 1 { //中序剩余一个元素,直接建树
		return root
	}
	index := getIndex(preorder[0], inorder)
	root.Left = buildTree1(preorder[1:index+1], inorder[:index]) //左子树,pre和in长度对应
	root.Right = buildTree1(preorder[index+1:], inorder[index+1:])
	return root
}

//从inOrder中找到每个根节点对应的索引
func getIndex(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}
