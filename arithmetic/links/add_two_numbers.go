package links

// AddTwoNumbers 2. 两数相加
//
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
// https://leetcode-cn.com/problems/add-two-numbers
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 思路：
// 可以从头节点开始，同时遍历两个链表，将数值加到其中一个链表上，另外需要一个额外变量，存储进一。
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/**
	if l1 == nil || l1.Val == 0 {
		return l2
	}
	if l2 == nil || l2.Val == 0 {
		return l1
	}

	res := l1

	var temp int
	for l1 != nil || l2 != nil {
		temp = l1.Val + l2.Val
		l1.Val = temp % 10
		temp = temp / 10

		if temp > 0 {
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1}
			} else {
				l1.Next.Val += 1
			}
		}

		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{}
		}
		if l2.Next == nil && l1.Next != nil {
			l2.Next = &ListNode{}
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	*/

	var tail, head *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
