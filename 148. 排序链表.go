package main

import (
	"fmt"
)

func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

//用两个指针p,q记录,都一直向前移动
//p之前的元素记录比基准值小的元素,这样每次遇到一个比基准值大的就和p的next指针的值交换
//最后为了让基准值在中间,再做一次交换
func partitionLinkList(head *ListNode, end *ListNode) *ListNode {
	num := head.Val
	p := head //p之前的元素都比num小
	q := p.Next
	for q != end {
		if q.Val < num {
			p = p.Next //p的下一个位置记录需要交换的位置
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	head.Val, p.Val = p.Val, head.Val
	return p
}
func quickSort(start *ListNode, end *ListNode) {
	if start != end {
		pivot := partitionLinkList(start, end)
		quickSort(start, pivot)
		quickSort(pivot.Next, end)
	}
}

func main() {
	ListNode1 := &ListNode{Val: 5}
	ListNode2 := &ListNode{Val: 19}
	ListNode3 := &ListNode{Val: 5}
	ListNode4 := &ListNode{Val: 3}
	ListNode1.Next = ListNode2
	ListNode2.Next = ListNode3
	ListNode3.Next = ListNode4
	ListNode4.Next = nil
	quickSort(ListNode1, nil)
	for ListNode1 != nil {
		fmt.Println(ListNode1.Val)
		ListNode1 = ListNode1.Next
	}
}
