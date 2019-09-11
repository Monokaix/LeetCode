package main

//思路：先序遍历，得到含有p q两个节点的路径
//比较两个路径，第一个不等的即为最近公共节点
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	path := []*TreeNode{}
	res1 := []*TreeNode{}
	res2 := []*TreeNode{}
	finish := false
	preorder(root, p, path, &res1, &finish)
	path = []*TreeNode{}
	finish = false
	preorder(root, q, path, &res2, &finish)
	length := 0
	//找到最短的
	if len(res1) < len(res2) {
		length = len(res1)
	} else {
		length = len(res2)
	}
	result := &TreeNode{}
	for i := 0; i < length; i++ {
		if res1[i] == res2[i] {
			result = res1[i] //这里记录的即是最后一个相同的节点
		}
	}
	return result
}

//先序遍历二叉树得到一条路径
//path 当前遍历的路径
//res 找到后最终返回的路径
//finish 标记，判断是否找到
func preorder(root *TreeNode, node *TreeNode, path []*TreeNode, res *[]*TreeNode, finish *bool) {
	if root == nil || *finish { //到达叶节点或者已经找到需要的node 这里是递归结束的条件
		return
	}
	path = append(path, root)
	if root == node {
		*finish = true
		*res = append(*res, path...)
	}
	preorder(root.Left, node, path, res, finish)
	preorder(root.Right, node, path, res, finish)
}
