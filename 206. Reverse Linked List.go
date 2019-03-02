package main



 type ListNode struct {
 	Val int
    Next *ListNode
}
// normal solution
//func reverseList(head *ListNode) *ListNode {
//	pre := &ListNode{}
//	pre = nil
//	cur := head
//	for cur != nil{
//		next := cur.Next
//
//		cur.Next = pre //reverse
//		pre = cur
//		cur = next
//	}
//	return pre
//}

// recursively
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
