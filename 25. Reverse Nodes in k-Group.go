package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	cur := head
	next := &ListNode{}

	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	head = dummyHead.Next
	//1->2->3 先反转成2->1->3 载反转成3->2-1,此时相当于2-1是一个整体
	for i := 0; i < len/k; i++ {
		// 一组k循环
		for j := 0; j < k-1; j++ {
			next = cur.Next
			cur.Next = next.Next

			next.Next = pre.Next
			pre.Next = next
		}
		// 更新进行下一组次反转
		pre = cur
		cur = cur.Next
	}
	return dummyHead.Next
}
