package main

func buildTree1(preorder []int, inorder []int) *TreeNode {
	//将中序对应的元素索引记录在map,方便直接在中序序列找到根
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return help(preorder, inorder, m)
}
func help(preorder []int, inorder []int, m map[int]int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(inorder) == 1 { //中序剩余一个元素,直接建树
		return root
	}
	index := m[preorder[0]]
	root.Left = buildTree1(preorder[1:index+1], inorder[:index]) //左子树,pre和in长度对应
	root.Right = buildTree1(preorder[index+1:], inorder[index+1:])
	return root
}
