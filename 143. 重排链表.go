package main

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre := &ListNode{}
	pre = nil
	cur := slow.Next //中点
	slow.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	//pre为反转后的头
	tmp := head
	for pre != nil {
		t1 := tmp.Next
		t2 := pre.Next
		tmp.Next = pre
		pre.Next = t1
		tmp = t1
		pre = t2
	}
}
