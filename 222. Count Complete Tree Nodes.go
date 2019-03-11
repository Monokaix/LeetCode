package main

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var h1 uint32 = 0
	var h2 uint32 = 0
	p := root
	q := root
	for p != nil {
		h1++
		p = p.Left
	}

	for q != nil {
		h2++
		q = q.Right
	}
	if h1 == h2 { // full binary tree
		return (1 << h1) - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
	// method 2
	/**
	if root == nil {
			return 0
		}
	return countNodes(root.Left) + countNodes(root.Right) + 1
	 */

}
