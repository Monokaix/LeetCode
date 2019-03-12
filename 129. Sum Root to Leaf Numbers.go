package main

func sumNumbers(root *TreeNode) int {
	return helper(root,0)
}
func helper(root *TreeNode,pre int) int{
	if root == nil{
		return 0
	}
	res := pre*10+root.Val
	if root.Left == nil && root.Right == nil{
		return res
	}else{
		return helper(root.Left,res)+helper(root.Right,res)
	}

}