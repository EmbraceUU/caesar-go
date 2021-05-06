package links

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head ListNode
	p := &head
	l := len(nums)
	for i, v := range nums {
		p.Val = v
		if i < l-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return &head
}
