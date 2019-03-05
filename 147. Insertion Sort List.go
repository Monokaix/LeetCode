package main

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	pre := dummyHead
	next := &ListNode{}
	if head == nil || head.Next == nil{
		return head
	}

	for cur != nil && cur.Next != nil{
		next = cur.Next
		//temp = temp.Next
		//cur = cur.Next

		// find position
		for pre.Next.Val < cur.Val{
			pre = pre.Next
		}
		// insert
		cur.Next = pre.Next
		pre.Next = cur
		pre = dummyHead
		cur = next
	}
	return dummyHead.Next
}
