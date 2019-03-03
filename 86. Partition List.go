package main

func partition(head *ListNode, x int) *ListNode {
	p1 := &ListNode{}
	p2 := &ListNode{}
	l := p1
	r := p2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		}else{
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = r.Next
	return l.Next
}
