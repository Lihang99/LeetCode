//利用快慢指针找到中点 进行链表翻转 双指针对比即可
func isPalindrome(head *ListNode) bool {
	if head ==nil{
		return false
	}
	mid := findMid(head)
	secondHalfHead := reverse(mid)
	p1,p2 :=head,secondHalfHead
	for p1!=nil&&p2!=nil{
		if p1.Val != p2.Val{
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func findMid(head *ListNode) *ListNode{
	slow,quick := head,head
	//快指针一次走两步 慢指针一次走一步 快指针走过头时慢指针就在中间
	for quick.Next !=nil && quick.Next.Next !=nil{
		quick = quick.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode{
	var prev *ListNode
	curr := head
	for curr != nil{
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}