package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead := &ListNode{0, nil}
	dummyHead.Next = head
	h := dummyHead
	for i := 0; i < m-1; i++ {
		h = h.Next
	}
	cur := h.Next
	next := cur.Next
	Nnext := next

	for i := 0; i < n-m; i++ {
		Nnext = next.Next
		next.Next = cur // reverse
		//继续向前移动
		cur = next
		next = Nnext
	}
	//衔接上
	h.Next.Next = Nnext
	h.Next = cur
	return dummyHead.Next

}
