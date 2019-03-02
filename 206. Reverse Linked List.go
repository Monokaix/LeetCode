package main



 type ListNode struct {
 	Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre = nil
	cur := head
	for cur != nil{
		next := cur.Next

		cur.Next = pre //reverse
		pre = cur
		cur = next
	}
	return pre
}
