package links

// removeNthFromEnd
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 思路：获取长度，借助【DummyNode】两次遍历
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// removeNthFromEndByStack 借助栈删除节点
func removeNthFromEndByStack(head *ListNode, n int) *ListNode {
	stack := make([]*ListNode, 0)
	dummy := &ListNode{0, head}
	cur := dummy
	for ; cur != nil; cur = cur.Next {
		stack = append(stack, cur)
	}
	// 0, 1, 2, 3, 4, 5
	prev := stack[len(stack)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// removeNthFromEndByDN 借助【双节点】，first和second节点之间相差 n 个节点，当first为末尾时，second为倒数第 n 个节点
func removeNthFromEndByDN(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first.Next != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}
