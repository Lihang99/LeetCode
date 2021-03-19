//本题没有给链表 不能使用常规方式写 因为被删除的不会是末尾节点 所以用这种巧妙的办法
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}