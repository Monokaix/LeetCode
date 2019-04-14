package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	cur := l3
	for l1 != nil && l2 !=nil{
		if l1.Val <= l2.Val{ //这里比较m+n次
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 == nil{
		cur.Next = l2
	}else{
		cur.Next = l1
	}

	return l3.Next
}
// 时间复杂度：最坏O（m+n） m、n分别为两个链表长度 平均O（min（m，n））
// 比如1，2，4和1，5比较
// 总共比较次数为（1，1）（2，1）（2，4）（3，4）共四次比较，最后把4再插入链表共m+n次操作