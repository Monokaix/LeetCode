package main

//分治,涉计递归
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n == 2 { //长度为2,直接调用两路归并算法,直接返回，其他情况递归调用
		return mergeTwoLists(lists[0], lists[1])
	}
	//拆分为两个链表递归调用
	mid := n / 2
	list1 := make([]*ListNode, 0)
	list2 := make([]*ListNode, 0)
	for i := 0; i < mid; i++ {
		list1 = append(list1, lists[i])
	}
	for i := mid; i < n; i++ {
		list2 = append(list2, lists[i])
	}
	l1 := mergeKLists(list1)
	l2 := mergeKLists(list2)

	return mergeTwoLists(l1, l2)
}
