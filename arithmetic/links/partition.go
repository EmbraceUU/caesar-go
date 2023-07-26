package links

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 将链表一分为二
	// 一个链表全部小于 x
	// 一个链表全部大于 x
	// 将两个链表合并
	small := new(ListNode)
	sp := small
	large := new(ListNode)
	lp := large

	for head != nil {
		if head.Val >= x {
			lp.Next = head
			lp = lp.Next
		} else {
			sp.Next = head
			sp = sp.Next
		}
		head = head.Next
	}

	sp.Next = large.Next
	lp.Next = nil
	return small.Next
}
