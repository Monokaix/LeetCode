package main

//O(1)的空间复杂度解决,O(n)的复杂度直接将链表元素进栈再依次弹栈和原链表元素比对即可
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//快慢指针找到链表中心元素
	slow := head
	fast := head
	//这里只用管fast指针后面为不为空即可,因为fast在slow后面,fast不走到nil,slow也走不到nil
	//同时这里注意边界,这样处理不管链表是奇数还是偶数个元素slow都可以到达中点进行判断
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next    //右半部分链表头
	new := reverse(right) //右半部分反转后的链表头
	for new != nil {
		if new.Val != head.Val {
			return false
		}
		new = new.Next
		head = head.Next
	}
	//最后再把链表反转回去以不破坏原来链表结构
	slow.Next = reverse(new)
	return true
}
func reverse(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre = nil
	cur := head
	for cur != nil {
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
