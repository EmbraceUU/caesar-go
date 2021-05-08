package linked_list

// DeleteDuplicates 给定一个按升序排序的链表
// 删除链表中重复的元素
// 错误记录：
// 1. 没有加上for循环
// 2. 没有加入当前节点p向后移动的逻辑，导致无限循环
// 3. 需要等重复节点全部删除完，再移到下一个节点
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func DeleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		if p.Val != p.Next.Val {
			// 前后值不一样时，往后移动
			p = p.Next
			continue
		}
		// 前后值一样时，删除后面的值，不移动p
		pn := p.Next.Next
		p.Next.Next = nil
		p.Next = pn
	}
	return head
}

// DeleteDuplicatesLatest DeleteDuplicates最终版
func DeleteDuplicatesLatest(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// DeleteDuplicatesII 给定一个按升序排序的链表，删除重复节点
// 1. 借助于dummy node来处理head
// 2. 借助于rmVal来删除所有重复节点
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func DeleteDuplicatesII(head *ListNode) *ListNode {
	// 使用dummy node来处理head
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	current := dummy
	// 链表里出现重复节点，最少有两个节点
	var rmVal int
	for current.Next != nil && current.Next.Next != nil {
		// 使用rmVal辅助变量来删除所有重复节点
		if current.Next.Val == current.Next.Next.Val {
			rmVal = current.Next.Val
			for current.Next != nil && current.Next.Val == rmVal {
				current.Next = current.Next.Next
			}
			continue
		}
		current = current.Next
	}
	return dummy.Next
}

// ReverseList 反转链表
// 使用一个fast和slow遍历list，并反转list
// 错误1. 没有加上slow.Next = nil，导致末尾出现无限循环
// https://leetcode-cn.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head.Next
	slow := head
	slow.Next = nil
	for fast != nil {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}

	return slow
}

// ReverseListLatest 反转链表，递归做法
// 关键是：递归到最后一个节点，作为新的头结点返回， 然后每一步都自己反转自己 head.Next.Next = head, 这样不需要借助其他节点。
func ReverseListLatest(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseListLatest(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
