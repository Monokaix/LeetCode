package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	preNode := l3
	flag := 0 //是否进位
	b1 := l1
	b2 := l2
	for b1 != nil || b2 != nil || flag > 0 {
		node := &ListNode{}
		val1 := 0
		val2 := 0
		if b1 != nil {
			val1 = b1.Val
			b1 = b1.Next
		}
		if b2 != nil {
			val2 = b2.Val
			b2 = b2.Next
		}
		node.Val = (val1 + val2 + flag) % 10
		flag = (val1 + val2 + flag) / 10

		preNode.Next = node
		preNode = node
	}
	return l3.Next
}
