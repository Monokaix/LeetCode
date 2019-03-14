package main

//双重递归
//以root为根的树中，查找sum，返回个数
func pathSum2(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath(root, sum)
	res += pathSum2(root.Left, sum)
	res += pathSum2(root.Right, sum)
	return res
}

//以node为根，查找
func findPath(node *TreeNode, num int) int {
	res := 0
	if node == nil {
		return 0
	}
	if node.Val == num {
		res += 1
	}
	res += findPath(node.Left, num-node.Val)
	res += findPath(node.Right, num-node.Val)
	return res
}
