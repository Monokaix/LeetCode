package main

//思路：递归
//后序遍历二叉树，运用递归，左右子树先展开好，再处理根节点
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root == nil{
		return
	}
	flatten(root.Left)
	flatten(root.Right) //此时左右子树已经处理好
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil{
		root = root.Right //找到左子树最后一个节点
	}
	root.Right = temp //找到后链接上右子树

}
