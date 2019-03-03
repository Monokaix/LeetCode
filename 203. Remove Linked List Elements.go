package main

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
