package main

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil && root.Val-sum == 0 {
		temp := []int{}
		temp = append(temp, root.Val)
		res = append(res, temp)
		return res
	}
	leftRes := pathSum(root.Left, sum-root.Val)
	for i := len(leftRes) - 1; i >= 0; i-- {
		temp := make([]int, 1)
		temp[0] = root.Val
		for _, v := range leftRes[i] {
			temp = append(temp, v)
		}

		leftRes[i] = temp
		res = append(res, leftRes[i])
	}
	rightRes := pathSum(root.Right, sum-root.Val)
	for i := len(rightRes) - 1; i >= 0; i-- {
		temp := make([]int, 1)
		temp[0] = root.Val
		for _, v := range rightRes[i] {
			temp = append(temp, v)
		}

		rightRes[i] = temp
		res = append(res, rightRes[i])
	}
	return res
}
