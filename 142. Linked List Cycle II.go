package main

//先判断是否有环,然后根据数学公式推倒分别从相遇处和头节点开始遍历,相遇即为交点节点
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil { //说明无环
		return nil
	}
	//处理有环的情况
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
