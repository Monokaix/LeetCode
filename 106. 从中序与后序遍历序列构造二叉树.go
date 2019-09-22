package main

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 { //中序剩余一个元素,直接建树
		return root
	}
	index := getIndex(postorder[len(postorder)-1], inorder)
	root.Left = buildTree2(inorder[:index], postorder[:index]) //左子树,pre和in长度对应
	root.Right = buildTree2(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

//从inOrder中找到每个根节点对应的索引
func getIndex2(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}
