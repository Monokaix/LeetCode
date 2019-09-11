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
	}
	//没找到解继续递归
	getAllPath(root.Left, sum-num, list, res)
	getAllPath(root.Right, sum-num, list, res)
	// list = list[:len(list)-1] 这里不需要回溯,因为append操作会产生一份新的拷贝,不影响当前list的值
}
