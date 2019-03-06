package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k== 0{
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	left := head
	right := dummyHead
	len := 0
	for  right.Next != nil {
		len++
		right = right.Next
	}
	k = k % len
	// reset right
	right = dummyHead
	for i:=0;i<k+1;i++{
		right = right.Next
	}

	for right.Next != nil{
		right = right.Next
		left = left.Next
	}
	right.Next = head
	dummyHead.Next = left.Next
	left.Next = nil
	return dummyHead.Next
}
