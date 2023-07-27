package links

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 两条路径 a b
	// a 遍历 headA, 然后遍历 headB
	// b 遍历 headB, 然后遍历 headA
	// 如果相交，最终会相遇，如果不相交，则返回 null

	a, b := headA, headB
	hasRangeB, hasRangeA := false, false
	for a != nil && b != nil {
		if a == b {
			return a
		}

		if a.Next == nil && !hasRangeB {
			hasRangeB = true
			a = headB
		} else {
			a = a.Next
		}

		if b.Next == nil && !hasRangeA {
			hasRangeA = true
			b = headA
		} else {
			b = b.Next
		}
	}
	return nil
}
