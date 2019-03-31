package main

import "math"

//对于一个和节点，有偷和不偷两种选择，计算偷或者不偷，然后取两者最大值
func rob4(root *TreeNode) int {
	m := make(map[*TreeNode]int, 0)
	return tryRob2(root, m)
}

func tryRob2(node *TreeNode, m map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	if _, ok := m[node]; ok {
		return m[node]
	}
	//偷取根节点
	res1 := 0
	if node.Left != nil {
		//从根节点的左节点的左节点开始偷
		res1 += tryRob2(node.Left.Left, m) + tryRob2(node.Left.Right, m)
	}
	if node.Right != nil {
		//从根节点的左节点的左节点开始偷
		res1 += tryRob2(node.Right.Left, m) + tryRob2(node.Right.Right, m)
	}
	res1 += node.Val
	//不偷根节点
	res2 := 0
	res2 += tryRob2(node.Left, m) + tryRob2(node.Right, m)
	m[node] = int(math.Max(float64(res1), float64(res2)))
	return m[node]
}
