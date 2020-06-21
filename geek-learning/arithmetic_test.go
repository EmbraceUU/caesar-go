package geek_learning

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	head := func() *ListNode {
		h := &ListNode{}
		next1 := &ListNode{Val: 1}
		h.Next = next1
		next2 := &ListNode{Val: 2}
		next1.Next = next2
		next3 := &ListNode{Val: 3}
		next2.Next = next3
		next4 := &ListNode{Val: 4}
		next3.Next = next4
		next5 := &ListNode{Val: 5}
		next4.Next = next5
		return h
	}()

	cur := head.Next
	println(head.Val)
	for {
		if cur == nil {
			break
		} else {
			println(cur.Val)
		}
		cur = cur.Next
	}

	head = ReverseLinkedListII(head)

	cur = head.Next
	println(head.Val)
	for {
		if cur == nil {
			break
		} else {
			println(cur.Val)
		}
		cur = cur.Next
	}
}
