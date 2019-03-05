package main

// four pointers
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	p := dummyHead
	for p.Next != nil && p.Next.Next != nil{
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next

		//swap
		node2.Next = node1
		node1.Next = next
		p.Next = node2

		// next traverse
		p = node1
	}

	return dummyHead.Next

}
