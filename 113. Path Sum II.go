package main

//典型的dfs,类似于数组中找路径,注意res要引用传递
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	list := make([]int, 0)
	getAllPath(root, sum, list, &res)
	return res
}
func getAllPath(root *TreeNode, sum int, list []int, res *[][]int) {
	if root == nil { //说明递归没有找到解,直接返回
		return
	}
	num := root.Val
	list = append(list, num)
	if root.Left == nil && root.Right == nil && sum == root.Val { //找到一个解,加入res
		tmp := make([]int, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
	} else { //没找到解继续递归
		if root.Left != nil {
			getAllPath(root.Left, sum-num, list, res)
		}
		if root.Right != nil {
			getAllPath(root.Right, sum-num, list, res)
		}
	}
}
