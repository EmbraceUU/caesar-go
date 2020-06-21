package geek_learning

/*
算法练习
*/

// reverse-linked-list
// 递归方式
func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := ReverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// 迭代方式
/*
在遍历列表时，将当前节点的 next 指针改为指向前一个元素。
由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。
在更改引用之前，还需要另一个指针来存储下一个节点。
不要忘记在最后返回新的头引用！
*/
func ReverseLinkedListII(head *ListNode) *ListNode {
	var prev *ListNode
	for {
		if head == nil {
			break
		}
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// swap-nodes-in-pairs
func SwapNodesInPairs() {

}

// linked-list-cycle
func LinkedListCycle() {

}

// linked-list-cycle-ii
func LinkedListCycleII() {

}

// reverse-nodes-in-k-group
func ReverseNodesInKCroup() {

}
