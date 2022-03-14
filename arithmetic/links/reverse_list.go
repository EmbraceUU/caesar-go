package links

// reverseList 反转链表
// 使用一个fast和slow遍历list，并反转list
// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1(s) <- 2(s) -> 3(t) -> 4 -> 5
	// 1(s) <- 2(s) <- 3(f) -> 4(s) -> 5(f)

	// 使用快慢指针，遍历过程中反转链表
	var f, s *ListNode = head, nil

	for f != nil {
		// 借助tmp节点保存下一个节点
		tmp := f.Next
		f.Next = s
		s = f
		f = tmp
	}

	return s
}
