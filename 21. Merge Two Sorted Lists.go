package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	cur := l3
	for l1 != nil && l2 !=nil{
		if l1.Val <= l2.Val{
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 == nil{
		cur.Next = l2
	}else{
		cur.Next = l1
	}

	return l3.Next
}
