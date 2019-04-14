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

	l1 := mergeKLists(lists[:mid])
	l2 := mergeKLists(lists[mid:])

	return mergeTwoLists(l1, l2)
}

// 方法一：笨方法：k个链表，前两个合并，然后合并第3个再合并第4个共合并k-1次，第一次比较最坏比较次数O（m+n）,
// 假设链表平均长度都为n，则为O（n+n），第二次为O（2n+n）,以此类推，相加得（n+n）+（2n+n）+（3n+n）+...
// +(k-1)n+n = (1+2+3+...+k-1)n+(k-1)n = k(k+1)/2*n-n=O(k^2*n)
// 方法二：k个链表，每个链表平均n个节点，两两合并，
// 第一轮进行k/2次，每次处理2n个数字
// 第二轮进行k/4次，每次处理4n个数字，以此类推
// 可以得到最后一轮共k/(2^logk)，处理2^logk*n个值
// 计算得到每一轮操作次数为kn，共进行logk次，所以时间复杂度为O(knlogk)
// 每次两两合并，合并完再两两合并最后得到一个完整得链表