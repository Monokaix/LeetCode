package main

//思路：二叉树递归，同时用快慢指针寻找中点
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return genBST(head, nil)
}

func genBST(head *ListNode, tail *ListNode) *TreeNode {
	if head == tail { //递归结束
		return nil
	}
	slow := head
	fast := head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//slow到达中心点
	root := &TreeNode{}
	root.Val = slow.Val
	root.Left = genBST(head, slow)
	root.Right = genBST(slow.Next, tail)
	return root
}
