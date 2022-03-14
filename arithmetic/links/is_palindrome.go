package links

// isPalindrome 请判断一个链表是否为回文链表。
// 借助【递归】
// https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	fp := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(node *ListNode) bool {
		if node != nil {
			if !recursivelyCheck(node.Next) {
				return false
			}
			if node.Val != fp.Val {
				return false
			}
			fp = fp.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

func isPalindromeByDN(head *ListNode) bool {
	if head == nil {
		return true
	}
	//找到前半部分链表的尾节点。
	firstHalfEnd := endOfFirstHalf(head)
	//反转后半部分链表。
	secondHalfStart := reverseList(firstHalfEnd.Next)
	//判断是否回文。
	p1, p2 := head, secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
			continue
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	//恢复链表。
	firstHalfEnd.Next = reverseList(secondHalfStart)
	//返回结果。
	return result
}

// endOfFirstHalf 查找链表的中点
func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
