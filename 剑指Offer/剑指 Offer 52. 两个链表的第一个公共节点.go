package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	lenA, lenB := 0, 0
	for ha != nil {
		lenA++
		ha = ha.Next
	}
	for hb != nil {
		lenB++
		hb = hb.Next
	}

	diff := lenB - lenA
	if lenA > lenB {
		diff = -diff
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < diff; i++ {
			headB = headB.Next
		}
	}

	for headB != headA {
		headB = headB.Next
		headA = headA.Next
	}

	return headA
}
