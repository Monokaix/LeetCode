package main

func buildTree2(inorder []int, postorder []int) *TreeNode {
	//将中序对应的元素索引记录在map,方便直接在中序序列找到根
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return help2(inorder, postorder, m)
}
func help2(inorder []int, postorder []int, m map[int]int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 { //中序剩余一个元素,直接建树
		return root
	}
	index := m[postorder[len(postorder)-1]]
	root.Left = buildTree2(inorder[:index], postorder[:index]) //左子树,pre和in长度对应
	root.Right = buildTree2(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
