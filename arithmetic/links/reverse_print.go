package links

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	// 可以入栈，然后逆序打印
	// 也可以使用递归的方式，后序遍历
	val := make([]int, 0)
	p(head, &val)
	return val
}

func p(head *ListNode, val *[]int) {
	if head == nil {
		return
	}

	p(head.Next, val)
	*val = append(*val, head.Val)
}
