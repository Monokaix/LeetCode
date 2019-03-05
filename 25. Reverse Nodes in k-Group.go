package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{}
	root.Next = head
	dummyHead := root
	temp := head

	len := 0
	for temp != nil{
		len++
		temp = temp.Next
	}

	for len >= k{
		// 一组k循环
		for i := 0 ;i< k-1;i++{
			temp = root.Next
			root.Next = head.Next
			head.Next = root.Next.Next
			root.Next.Next = temp
		}
		// 更新进行下一组次反转
		root = head
		head = head.Next
		len -= k
	}
	return dummyHead.Next

}