package main



 type ListNode struct {
 	Val int
    Next *ListNode
}
// normal solution
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 2a48fd708ab7a7f03944223346c8bf863e1d41f4
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
<<<<<<< HEAD
>>>>>>> 2a48fd708ab7a7f03944223346c8bf863e1d41f4
=======
>>>>>>> 2a48fd708ab7a7f03944223346c8bf863e1d41f4
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// recursively
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil{
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//}
