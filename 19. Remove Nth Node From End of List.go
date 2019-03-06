package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	p := dummyHead
	q := dummyHead
	for i:=0;i<n+1;i++{
		q = q.Next
	}
	for q != nil{
		p = p.Next
		q = q.Next
	}
	// delete
	p.Next = p.Next.Next
	return dummyHead.Next

}
