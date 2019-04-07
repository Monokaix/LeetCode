package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	top1 := 0
	stack2 := make([]int, 0)
	top2 := 0
	//链表元素入栈
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		top1++
		l1 = l1.Next
	}
	//链表元素入栈
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		top2++
		l2 = l2.Next
	}
	flag := 0
	dummyHead := &ListNode{}
	for top1 > 0 || top2 > 0 || flag > 0 {
		node := &ListNode{}
		val1 := 0
		val2 := 0
		if top1 > 0 {
			top1--
			val1 = stack1[top1]
			stack1 = stack1[:top1]
		}
		if top2 > 0 {
			top2--
			val2 = stack2[top2]
			stack2 = stack2[:top2]
		}
		node.Val = (val1 + val2 + flag) % 10
		flag = (val1 + val2 + flag) / 10
		node.Next = dummyHead.Next //注意这里是头插法
		dummyHead.Next = node
	}
	return dummyHead.Next
}
