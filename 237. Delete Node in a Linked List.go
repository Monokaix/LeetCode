package main

func deleteNode(node *ListNode) {
	if node == nil{
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
