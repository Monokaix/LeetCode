package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	p1 := head
	p2 := p1.Next
	odd := p1
	even := p2
	for p1.Next != nil && p2.Next != nil{
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = even
	return odd
}

