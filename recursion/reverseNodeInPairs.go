package recursion

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//     first := head
	//     second := head.Next

	//     first.Next = swapPairs(second.Next)
	//     second.Next = first
	//     return second
	newHead := head.Next
	temp := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairs(temp)
	return newHead
}
