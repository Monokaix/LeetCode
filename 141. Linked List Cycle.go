package main

import "fmt"

//快慢指针法
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func main() {
	node1 := &ListNode{
		1,
		nil,
	}
	node2 := &ListNode{
		2,
		nil,
	}
	node1.Next = node2
	fmt.Println(hasCycle(node1))
}
