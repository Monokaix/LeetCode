package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res1 := []*TreeNode{}
	res2 := []*TreeNode{}
	path := []*TreeNode{}
	find := false

	findPath(root, p, path, &find, &res1)
	find = false
	findPath(root, q, path, &find, &res2)

	minLen := 0
	if len(res1) < len(res2) {
		minLen = len(res1)
	} else {
		minLen = len(res2)
	}

	res := &TreeNode{}
	for i := 0; i < minLen; i++ {
		if res1[i] == res2[i] {
			res = res1[i]
		} else {
			break
		}
	}

	return res
}

func findPath(root, target *TreeNode, path []*TreeNode, find *bool, res *[]*TreeNode) {
	if root == nil || *find {
		return
	}

	path = append(path, root)
	if root == target {
		*find = true
		// 这里path不是引用传递，正是为了让path只存单一路径，而不是整个遍历的路径
		*res = append(*res, path...)
	}

	findPath(root.Left, target, path, find, res)
	findPath(root.Right, target, path, find, res)
}
