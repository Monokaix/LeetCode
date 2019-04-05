package main

// 方法一:利用一个set,将链表A地址放入set,再将链表B地址放入set,第一个重复元素即为交点
// 时间O(n),空间O(n)

// 方法二:求出两个链表的长度差a-b,先让较长的链表移动a-b的长度,两个链表指针对齐，
// 然后一起遍历，找到相同的即为交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := 0
	b := 0
	cura := headA
	curb := headB
	for cura != nil {
		a++
		cura = cura.Next
	}
	for curb != nil {
		b++
		curb = curb.Next
	}
	cura = headA
	curb = headB
	if a > b {
		for i := 0; i < a-b; i++ {
			cura = cura.Next
		}
	} else {
		for i := 0; i < b-a; i++ {
			curb = curb.Next
		}
	}
	//指针起始位置一样,开始同步遍历
	for cura != nil {
		if cura == curb {
			break
		}
		cura = cura.Next
		curb = curb.Next
	}
	return cura
}
